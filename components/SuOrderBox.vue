<template>
  <div class="card bg-box p-4 p-3 mb-3 mb-md-0">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-18">Order Breakdown</p>
      </div>
    </div>
    <div class="row align-items-center mb-lg-3 mt-2">
      <div class="col-6">
        <h1 class="text-white">69</h1>
      </div>
      <div class="col-6 text-right order-lg-last">
        <span class="ff-poppins-sb fs-18">Orders</span>
      </div>
    </div>

    <div class="row align-items-center mb-1">
      <div class="col-6">
        <h1 class="text-white">$69.69</h1>
      </div>
      <div class="col-6 text-right order-last">
        <span class="ff-poppins-sb fs-18">Average Order Value</span>
      </div>
    </div>

  </div>
</template>

<script>
import SuInfoItem from "@/components/SuInfoItem";
export default {
  name: 'SuOrderBox',
  components: {SuInfoItem},
  props: {
    hasMore: {
      type: Boolean,
      default: false
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
