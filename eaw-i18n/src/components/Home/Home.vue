<template>
  <div class="hello">
    <a v-bind:href="website">GOOGLE</a>
    <input type="text" v-model="name" v-on:keyup="nameUpdated" />
    <input type="text" v-model="name2" v-on:keyup="name2Updated" />

    <div>Your name is {{ name }}</div>

    <md-button class="md-raised">test</md-button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import axios from 'axios';

@Component
export default class Home extends Vue {
  name = 'Random';
  name2 = 'Random2';
  website = 'http://google.com';
  a = 2;

  mounted(): void {
    console.log('mounted');

    this.getData().then(resp => {
      console.log('Got data');
    });
  }

  getData(): Promise<any> {
    return axios.get('http://localhost:9090').then(resp => {
      console.log(resp);
    });
  }

  nameUpdated() {
    console.log('name', this.name);

    this.name2 += 'ha';
  }

  name2Updated() {
    console.log('name2', this.name2);
  }
}
</script>
<style scoped lang="scss">
$test: 30px;

h3 {
  margin: $test 0 0;
}
</style>
