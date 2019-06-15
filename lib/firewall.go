package lib

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"sort"
	"strings"
)

func ServeReverseProxy(config WafConfig, res http.ResponseWriter, req *http.Request) {

	// Filter requests by IP address.
	ip := getClientIP(req)
	if !ipIsAllowed(ip, config.IpFilterMode, config.IpAddresses) {
		log.Printf("Request from IP %s was blocked.", ip)
		blockRequest(res)
		return
	}


	// Test for injection attacks.
	if requestContainsInjection(req) {
		blockRequest(res)
		return
	}

	if strings.HasPrefix(req.Header.Get("Content-Type"), "multipart/form-data") {
		reqBody, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		//Check if this is a file upload request.
		err := req.ParseMultipartForm(5*1024*1024)
		if  err == nil {

			// Prepare file extensions in a regular expression.
			blockExtensions := strings.Join(config.DenyExtensions, "|")
			expr, _ := regexp.Compile("\\.(" + blockExtensions + ")$")

			// Go through all uploaded files.
			for key, files := range req.MultipartForm.File {
				for _, file := range files {
					// Check if the filename against the regular expression.
					if expr.MatchString(file.Filename) {
						log.Printf("File upload in %s with filename %s was blocked.", key, file.Filename)
						blockRequest(res)
						return
					}
				}
			}
		} else {
			log.Println(err)
		}
		// Re-assign the body in the request after the inspection.
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
	}


	// All checks passed, prepare the proxy.
	targetUrl, _ := url.Parse(config.Upstream)
	director := func(req *http.Request) {
		// Update some headers to allow for HTTPS redirection.
		req.URL.Host = targetUrl.Host
		req.URL.Scheme = targetUrl.Scheme
	}

	proxy := httputil.ReverseProxy{Director: director}

	log.Printf("Forwarding request to %s\n", config.Upstream)
	proxy.ServeHTTP(res, req)

	// Remove sensitive information.
	res.Header().Del("Server")
}

func blockRequest(res http.ResponseWriter) {
	res.WriteHeader(http.StatusNotAcceptable)
}

func requestContainsInjection(req *http.Request) bool {
	// Check URL query params.
	for _, param := range req.URL.Query() {
		for _, value := range param {
			if stringContainsInjection(value) {
				return true
			}
		}
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		return false
	}

	// Check form params.
	if  req.ParseForm() == nil {
		for _, values := range req.PostForm {
			for _, value := range values {
				if stringContainsInjection(value) {
					return true
				}
			}
		}
	}

	return false
}

func stringContainsInjection(str string) bool {
	if hasSQLi, sig := TestSQLi(str); hasSQLi {
		log.Printf("Found SQL injection %s in %s", sig, str)
		return true
	}

	if TestXSS(str) {
		log.Printf("Found XSS %s", str)
		return true
	}

	return false
}

func getClientIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return IPAddress
}

func ipIsAllowed(ip string, mode string, ipList []string) bool {
	// ipList is already sorted when parsing the config.
	ipMatches := sort.SearchStrings(ipList, ip)

	switch mode {
	case "blacklist":
		// IP is not found in the blacklist.
		if ipMatches == len(ipList) {
			return true
		}
	case "whitelist":
		// IP is found in the whitelist.
		if ipMatches < len(ipList) {
			return true
		}

	}

	// Block by default.
	return false
}