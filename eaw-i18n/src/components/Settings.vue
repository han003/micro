<template>
  <div>
    <span>this is settings</span>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class Settings extends Vue {
  mounted(): void {
    console.log('Settings');
  }
}
</script>
<style scoped lang="scss">
</style>
