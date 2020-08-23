<template>
  <div>
    <span>This is home</span>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class Home extends Vue {
  mounted(): void {
    console.log('home');
  }
}
</script>

<style scoped lang="scss">
</style>
