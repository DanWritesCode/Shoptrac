<template>
  <div>
    <navbar></navbar>
    <nuxt />
    <v-footer></v-footer>
  </div>
</template>
<script>
import Navbar from '../components/navbar'
import VFooter from '../components/v-footer'
export default {
  components: { VFooter, Navbar }
}
</script>

<style lang="scss">
@import '../assets/sass/variables';
body {
  font-family: 'Noto Sans', sans-serif;
  color: $text;
  font-size: 14px;
}
</style>
