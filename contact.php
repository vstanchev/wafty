<?php
    if (!empty($_REQUEST)) {
        foreach ($_REQUEST as $key => $value) {
            if (strpos($value, 'OR') !== false || strpos($value, 'script') !== false) {
                header('Not acceptable', true, 406);
                echo '406 Not acceptable';
                exit();
            }
        }
    }
    if (!empty($_FILES)) {
        foreach ($_FILES as $key => $value) {
            if (strpos($value['name'], '.exe') !== false || strpos($value['name'], '.php') !== false) {
                header('Not acceptable', true, 406);
                echo '406 Not acceptable';
                exit();
            }
        }
        echo '<h3>Заявката е приета. Ще се свържем с вас възможно най-скоро.</h3>';
        exit();
    }
?>

<html>
<head>
  <style>
    label {
      display: block;
    }
    textarea {
      display: block;
    }
  </style>
</head>
<body>

<h1>Запитване за изработка по поръчка</h1>
<form action="" method="POST" enctype="multipart/form-data">
    <label for="name">Вашето име:</label>
    <input type="text" name="name"/>

    <label for="name">Проект:</label>
    <input type="file" name="project"/>

    <label for="comment">Коментар:</label>
    <textarea cols=50 rows=5 name="comment"></textarea>

    <input type="submit" value="Изпрати">
</form>
</body>
</html>
