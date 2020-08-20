<template>
  <div class="hello">
    <a v-bind:href="website">GOOGLE</a>
    <input type="text" v-model="name" v-on:keyup="nameUpdated" />
    <input type="text" v-model="name2" v-on:keyup="name2Updated" />

    <div>Your name is {{ name }}</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class HelloWorld extends Vue {
  name = 'Random';
  name2 = 'Random2';
  website = 'http://google.com';
  a = 2;

  nameUpdated() {
    console.log('name', this.name);

    this.name2 += 'ha';
  }

  name2Updated() {
    console.log('name2', this.name2);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
