<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html xmlns:og="http://ogp.me/ns#">
<head>
  <title>Daozai Kafka</title>
  <link rel='stylesheet' href='/static/css/styles.css' type='text/css'>
</head>

<body>
<div class="main">
  <div class="header">
    <a href="/"><img width="325" height="97" class="logo" src="../static/images/logo.png"></a>
  </div>

  <div class="content">
    <nav class="b-sticky-nav">
      <div class="content__left">
        <div class="nav__inner">
          <a class="nav__item" href="/"><h1>kafka</h1></a>
          <a class="nav__item" href="/kafka/topiclist">all topic</a>
          <a class="nav__item" href="/kafka/topicinfo">topic info</a>
          <a class="nav__item" href="/kafka/createtopic">create topic</a>
          <a class="nav__item" href="/kafka/grouplist">all group</a>
          <a class="nav__item" href="/kafka/groupinfo">group info</a>
          <a class="nav__item" href="kafka/alterpartition">alter partition</a>
        </div>
      </div>
    </nav>

    <div class="content__middle">
      <form action="/kafka/createtopic" method="post">
        zookeeper:<input class="content_middle_sub" name="zookeeper"><br>
        topic:<input class="content_middle_sub" name="topic"><br>
        partition:<input class="content_middle_sub" name="partition"><br>
        replicationfactor:<input class="content_middle_sub" name="replicationfactor"><br>
        <input type="submit" value="执行">
      </form>


    </div>
    <div class="content__right">


    </div>
  </div>
</div>
<div class="footer">
  <div class="footer__inner">
    <div class="footer__legal">
    </div>
  </div>
</div>
</body>
</html>
