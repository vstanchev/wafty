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
    .comment {
      width: 400px;
    }
  </style>
</head>
<body>

<h1>Продукт 1</h1>
  <div style="float:left">
    <img style="width: 400px" src="https://www.witneyandmason.com/wp-content/themes/witneyandmason/images/product-placeholder.gif">
  </div>

  <div style="float:left; width: 500px">
    <h3>Продукт 1</h3>
    <strong>52,33 лв</strong>
    <p>Lorem Ipsum е елементарен примерен текст, използван в печатарската и типографската индустрия. Lorem Ipsum е индустриален стандарт от около 1500 година, когато неизвестен печатар взема няколко печатарски букви и ги разбърква, за да напечата с тях книга с примерни шрифтове. </p>
    <form action="" method="POST">
      <label for="name">Вашето име:</label>
      <input type="text" name="name"/>
      <label for="comment">Коментар:</label>
      <textarea cols=50 rows=5 name="comment"></textarea>
      <input type="submit" value="Изпрати">
    </form>
  </div>
  <div style="clear:both"></div>
  <h2>Коментари</h2>
  <?php if (isset($_REQUEST['name']) && isset($_REQUEST['comment'])): ?>
    <div class="comment">
      <em><?=$_REQUEST['name']?> (07.07.2019)</em>
      <blockquote><?=$_REQUEST['comment']?></blockquote>
      <hr>
    </div>
  <?php endif; ?>
  <div class="comment">
    <em>Иван Иванов (07.07.2019)</em>
    <blockquote>Lorem Ipsum е елементарен примерен текст, използван в печатарската и типографската индустрия.</blockquote>
    <hr>
  </div>
  <div class="comment">
    <em>Георги Георгиев (01.03.2019)</em>
    <blockquote>Lorem Ipsum е елементарен примерен текст, използван в печатарската и типографската индустрия.</blockquote>
    <hr>
  </div>
</body>
</html>
