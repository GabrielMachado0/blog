<template>
<body>
  <div class="container-post">
    <img src="./adp-logo.png" alt="" />
    <div class="bt-login">
      <input type="email" name="email" placeholder="Seu e-mail..." />
      <input type="password" name="senha" placeholder="Sua senha..." />
      <a href="./logado"><input type="submit" value="Conectar" /></a>
      <a href="./cadastro"><p class="cadastro">Cadastrar, clique aqui</p></a>
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
  margin-bottom: -5%;
}

.cadastro {
  font-size: 9px;
  color: white;
  margin: 0 auto;
  margin-top: -7px;
  margin-left: 35px;
}

.cadastro:hover {
  color: #ca652a;
}

.bt-login input {
  display: inline-block;
  margin: 10px auto;
  margin-left: 10px;
}

.bt-login {
  margin-left: 48%;
  margin-top: 10px;
}

.bt-login input[type=email],
.bt-login input[type=password] {
  padding-left: 9px;
  padding-right: 7px;
  border-radius: 20px;
  font-size: 13px;
  max-width: 140px;
  width: 100%;
  height: 25px;
  border: 1px solid #ccc;
}

.bt-login input[type=submit] {
  border-radius: 20px;
  font-size: 13px;
  font-weight: bold;
  max-width: 80px;
  width: 100%;
  cursor: pointer;
  color: white;
  height: 25px;
  background-color: #923905;
}

.bt-login input[type=submit]:hover {
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
