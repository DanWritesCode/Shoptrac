<template>
  <div class="card bg-box p-lg-4 p-3 h-100">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-18">{{title}}</p>
      </div>
    </div>

    <div class="row align-items-center mt-2">
      <div class="col-5">
        <span class="ls-2 ff-poppins-sb">Name</span>
      </div>
      <div class="d-none d-bp d-xl-block col-lg-4">
        <span class="ls-2 ff-poppins-sb">
          Units Sold
        </span>
      </div>
      <div class="col-lg-3 col-md-2 d-none d-lg-block">
        <span class="ls-2 ff-poppins-sb">Total Revenue</span>
      </div>
    </div>
    <div class="mb-4" v-for="(key, val) in data">
      <su-info-item :name="key" :val="val" status="is-blue">
      </su-info-item>
    </div>
  </div>
</template>

<script>
import SuInfoItem from "@/components/SuInfoItem";
export default {
  name: 'SuListBox',
  components: {SuInfoItem},
  props: {
    title: {
      type: String,
      default: ""
    },
    data: {
      type: Array,
      default: {}
    }
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
