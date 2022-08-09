<template>
  <div class="form">
    <img src="./adp-logo.png" alt="">
    <input type="text" name="titulo" placeholder="Seu titulo..." />
    <textarea placeholder="Sua mensagem aqui..."></textarea>
    <input type="submit" value="Enviar" />
  </div>
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

.form {
  text-align: center;
  padding: 20px;
}

.form img {
  max-width: 500px;
  width: 100%;
}

.form input[type=text] {
  padding-left: 10px;
  display: block;
  border-radius: 20px;
  margin: 10px auto;
  font-size: 16px;
  max-width: 500px;
  width: 100%;
  height: 35px;
  border: 1px solid #ccc;
}

.form textarea {
    padding: 10px;
    resize: none;
    display: block;
    border-radius: 20px;
    margin: 10px auto;
    font-size: 16px;
    max-width: 500px;
    width: 100%;
    height: 120px;
}

.form input[type=submit] {
  padding-left: 10px;
  display: block;
  border-radius: 20px;
  margin: 10px auto;
  font-size: 16px;
  font-weight: bold;
  max-width: 500px;
  width: 100%;
  cursor: pointer;
  color: white;
  height: 35px;
  background-color: #923905;
}

.form input[type=submit]:hover {
  background-color: #ca652a;
}
</style>

<script>
export default {

}

  /*created() {
    this.carregarArtigos();
  
  },*/

  /*methods: {
    carregarArtigos() {
      this.$axios
        .$get("")
        .then((respostaAPI) => {
          this.listaArtigos = respostaAPI.data;
        })
    }
  }*/

</script>
