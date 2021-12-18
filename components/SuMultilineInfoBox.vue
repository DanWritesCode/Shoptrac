<template>
  <div class="card bg-box p-4 p-3 mb-3 mb-md-0">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-20">{{title}}</p>
      </div>
    </div>

    <div class="row align-items-center mb-1" v-for="(key, val) in data">
      <div class="col-6" >
        <h1 :class="color">{{key}}</h1>
      </div>
      <div class="col-6 text-right order-last">
        <span class="ff-poppins-sb fs-18">{{val}}</span>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'SuMultilineInfoBox',
  components: {},
  props: {
    title: {
      type: String,
      default: ""
    },
    data: {
      type: Array,
      default: {}
    },
    color: {
      type: String,
      default: "text-white"
    },
  }
}
</script>

<style scoped>
.content-box {
  height: 380px;
  overflow-y: auto;
  overflow-x: hidden;
}
</style>
