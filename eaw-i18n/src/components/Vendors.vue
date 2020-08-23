<template>
  <div>
    <span>this is vendors</span>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class Vendors extends Vue {
  mounted(): void {
    console.log('venors');
  }
}
</script>
<style scoped lang="scss">
</style>
