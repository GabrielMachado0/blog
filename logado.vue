<template>
<body>
  <div class="container-post">
    <img src="./adp-logo.png" alt="" />
    <div class="bt-login">
      <p class="nome-user">Gabriel Machado</p>
        <a href="/"><p class="logout">Sair</p></a>
    </div>
    <article class="container-text">
      <h2 class="titulo-artigo">Lorem Ipsum</h2>
        <span class="conteudo">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce laoreet lectus ac ex congue consectetur. 
          Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Morbi aliquet venenatis diam, id ultrices lectus suscipit id. 
          Etiam at aliquam libero. Aliquam maximus tincidunt sem, ut consectetur nunc. In sed elit iaculis, finibus nibh et, consectetur sem. Aenean id tincidunt orci. 
          Donec vehicula, sapien a tincidunt porta, nibh dolor porttitor risus, sit amet lacinia lorem orci eu nibh. Phasellus sed facilisis augue.          
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed bibendum ipsum ac bibendum pellentesque. Nam ut ullamcorper urna. Proin id purus mollis, 
          placerat nulla eget, faucibus urna.
        </span>
        <div class="container-bt">
          <a href=""><button class="bt-leia">Leia Mais</button></a>
        </div>
    </article>
    <article class="container-text">
      <h2 class="titulo-artigo">Lorem Ipsum</h2>
        <span class="conteudo">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce laoreet lectus ac ex congue consectetur. 
          Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Morbi aliquet venenatis diam, id ultrices lectus suscipit id. 
          Etiam at aliquam libero. Aliquam maximus tincidunt sem, ut consectetur nunc. In sed elit iaculis, finibus nibh et, consectetur sem. Aenean id tincidunt orci. 
          Donec vehicula, sapien a tincidunt porta, nibh dolor porttitor risus, sit amet lacinia lorem orci eu nibh. Phasellus sed facilisis augue.          
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed bibendum ipsum ac bibendum pellentesque. Nam ut ullamcorper urna. Proin id purus mollis, 
          placerat nulla eget, faucibus urna.
        </span>
        <div class="container-bt">
          <a href=""><button class="bt-leia">Leia Mais</button></a>
        </div>
    </article>
    <article class="container-text">
      <h2 class="titulo-artigo">Lorem Ipsum</h2>
        <span class="conteudo">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce laoreet lectus ac ex congue consectetur. 
          Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Morbi aliquet venenatis diam, id ultrices lectus suscipit id. 
          Etiam at aliquam libero. Aliquam maximus tincidunt sem, ut consectetur nunc. In sed elit iaculis, finibus nibh et, consectetur sem. Aenean id tincidunt orci. 
          Donec vehicula, sapien a tincidunt porta, nibh dolor porttitor risus, sit amet lacinia lorem orci eu nibh. Phasellus sed facilisis augue.          
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed bibendum ipsum ac bibendum pellentesque. Nam ut ullamcorper urna. Proin id purus mollis, 
          placerat nulla eget, faucibus urna.
        </span>
        <div class="container-bt">
          <a href=""><button class="bt-leia">Leia Mais</button></a>
        </div>
    </article>
    <article class="container-text">
      <h2 class="titulo-artigo">Lorem Ipsum</h2>
        <span class="conteudo">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce laoreet lectus ac ex congue consectetur. 
          Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Morbi aliquet venenatis diam, id ultrices lectus suscipit id. 
          Etiam at aliquam libero. Aliquam maximus tincidunt sem, ut consectetur nunc. In sed elit iaculis, finibus nibh et, consectetur sem. Aenean id tincidunt orci. 
          Donec vehicula, sapien a tincidunt porta, nibh dolor porttitor risus, sit amet lacinia lorem orci eu nibh. Phasellus sed facilisis augue.          
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed bibendum ipsum ac bibendum pellentesque. Nam ut ullamcorper urna. Proin id purus mollis, 
          placerat nulla eget, faucibus urna.
        </span>
        <div class="container-bt">
          <a href=""><button class="bt-leia">Leia Mais</button></a>
        </div>
    </article>
  </div>
</body>
</template>

<style>
* {
  margin: 0 auto;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Lato', sans-serif;
}

body {
  background-color: #27282D;
}

div .container-post img {
  margin-right: 61%;
  margin-bottom: -4.5%;
}

.bt-login {
  display: inline-block;
  margin: 10px auto;
  margin-left: 65%;
  
}

.nome-user {
  margin-top: 2%;
  margin-bottom: -10%;
  padding-left: 7px;
  padding-right: 7px;
  border-radius: 20px;
  font-size: 13px;
  max-width: 140px;
  width: 100%;
  height: 22px;
  border: 1px solid #ccc;
  color: white;
  background-color: #923905;
}

.logout {
  margin-top: 12%;
  margin-left: 33%;
  margin-bottom: -10%;
  padding-left: 7px;
  padding-right: 7px;
  border-radius: 20px;
  text-align: center;
  font-size: 13px;
  max-width: 40px;
  width: 100%;
  height: 22px;
  border: 1px solid #ccc;
  color: white;
  background-color: #923905;
}

.logout:hover {
  background-color: #ca652a;
}

.container-post {
  text-align: left;
  padding: 20px;
}

article {
  background-color: rgba(0, 162, 255, 0.336);
  border-radius: 20px;
  box-shadow: 0 0 1em black;
  padding: 0 7px;
  width: 45%;
  margin:  0 auto;
  margin-top: 40px;
}

.titulo-artigo {
  color: white;
  font-size: 30px;
  font-weight: bold;
  margin-bottom: 2%;
}

.conteudo {
  text-align: left;
  padding: 20px;
  color: white;
  font-size: 15px;
  font-weight: 400;
}

.bt-leia {
  margin-top: 15px;
  margin-bottom: 10px;
  padding: 5px;
  background-color: #923905;;
  border-radius: 20px;
  color: white;
  font-weight: bold;
  font-size: 15px;
}

.bt-leia:hover {
  background-color: #ca652a;
}

</style>

<script>
export default {

}

</script>
