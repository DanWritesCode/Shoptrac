<template>
  <div id="app">
    <div id="content">
        <div class="container-fluid p-4 h-100">
          <nuxt/>
        </div>
    </div>
  </div>
</template>
<script>
import VFooter from '../components/v-footer'

export default {
  components: {VFooter}
}
</script>
<style lang="scss">
@import url('https://fonts.googleapis.com/css?family=Poppins:200,300,400,600,700,800');
@import '../assets/sass/variables';

body {
  font-family: 'Poppins', sans-serif;
  color: $text;
  font-size: 14px;
  background: $bg;
}
#app {
  display: flex;
  align-items: stretch;
  position: relative;
}

#content {
  min-height: 100vh;
  width: 100vw;
  transition: all 0.3s;
}

#wrapper {
  padding: 10px;
}
</style>
