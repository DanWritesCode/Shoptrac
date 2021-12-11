<template>
  <div class="card bg-box p-lg-4 p-3 h-100">
    <div class="row text-center">
      <div class="mx-auto">
        <h5>Order Breakdown</h5>
      </div>
    </div>
    <div class="row align-items-center mb-lg-3 mt-2">
      <div class="col-lg-8 d-none d-lg-block">
        <h1 class="text-white">69</h1>
      </div>
      <div class="col-12 col-lg-4 text-lg-right text-left order-lg-last order-first">
        <span class="ff-poppins-sb fs-18">Orders</span>
      </div>
    </div>

    <div class="row align-items-center mb-lg-1">
      <div class="col-lg-8 d-none d-lg-block">
        <h1 class="text-white">$69.69</h1>
      </div>
      <div class="col-12 col-lg-4 text-lg-right text-left order-lg-last order-first">
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
