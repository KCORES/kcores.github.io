// reading-template.go
package main

const READING_P1 = `<!DOCTYPE html>
<html>

<head>
  <!-- Global site tag (gtag.js) - Google Analytics -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=UA-180293201-1"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag() { dataLayer.push(arguments); }
    gtag('js', new Date());

    gtag('config', 'UA-180293201-1');
  </script>

  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <title> KCORES - 氪金核心 </title>
  <link rel="shortcut icon" href="assets/logo/kcores-logo.ico" />
  <meta name="keywords" content=" KCORES, 氪金核心, 消费电子, 万兆网络, 家用NAS, 垃圾佬, 服务器, 云服务及主机, 仪表, Homelab, Vintage" />
  <meta name="description" content="KCORES - 氪金核心. 一个奇特的电子产品爱好者网站. 该项目由 @karminski-牙医 发起, 目的是建立一个大家能轻松讨论泛计算机话题的环境." />
  <meta name="viewport" content="width=device-width, user-scalable=no">
  <!-- include scripts & styles -->
  <link rel="stylesheet" type="text/css" href="assets/styles/rem.css">
  <link rel="stylesheet" type="text/css" href="assets/styles/reset.css">
  <link rel="stylesheet" type="text/css" href="assets/styles/base.css">
  <link rel="stylesheet" type="text/css" href="assets/styles/reading.css">
  <script src="assets/js/minigrid.js"></script>
  <script src="https://unpkg.com/imagesloaded@4/imagesloaded.pkgd.min.js"></script>
  

</head>

<body>

  <div id="main">

    <!-- Herader -->
    <header>
      <div class="inner ">
        <nav class="left">
          <a href="https://kcores.com/" class="logo">
            <h1>KCORES 氪金核心</h1>
          </a>
          <a href="https://kcores.com/reading" class="nav-button">
            <h1>阅读</h1>
          </a>
          <a href="https://kcores.com/topics" class="nav-button">
            <h1>话题</h1>
          </a>
        </nav>
        <nav class="right">
          <a href="https://kcores.com/about">关于</a>
          <a href="https://kcores.com/blogroll">友情链接</a>
          <a target="_blank" href="https://github.com/KCORES">Github</a>
        </nav>
      </div>
    </header>
  </div>


  <!-- grid content -->
  <div class="water-basic">
`


const READING_P2 = `
    <!-- THIS CONTENT IS AUTOMATIC GENERATED BY CONTENT-BUILDER, DO NOT EDIT THIS FILE MANUALLY -->
    <div class="item">
      <div class="entry-image box%d">
        <img src="%s" alt="">
      </div>
      <div class="entry-title"><a target="_blank" href="%s">%s</a></div>
      <div class="entry-info">%s / %s</div>
    </div>
`

const READING_P3 = `
  </div>

  <!-- load js -->

  <script>

    (function () {
      var imgLoad = imagesLoaded('.water-basic', function () {
      })
      imgLoad.on('progress', function (instance, image) {
        console.log('image: ', image);
        image.img.style.minHeight='auto'
        console.log('instance: ', instance);
        var grid;
        function init() {
          grid = new Minigrid({
            container: '.water-basic',
            item: '.item',
            gutter: 12,
            animate: function () {
              console.log(222)
            }
          });
          grid.mount();
        }
        // mount
        function update() {
          setTimeout(function () {
            grid.mount();

          }, 100)
        }
        init()
        window.addEventListener('resize', update);
      })
    })()

  </script>
</body>

</html>

`