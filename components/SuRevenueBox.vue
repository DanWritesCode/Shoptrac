<template>
  <div class="card bg-box p-4 p-3 mb-5 mb-md-0 h-100">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-18">Revenue Breakdown</p>
      </div>
    </div>
    <div class="row align-items-center mb-lg-3 mt-2">
      <div class="col-6">
        <h1 class="text-white">$9876.54</h1>
      </div>
      <div class="col-6 text-right order-last">
        <span class="ff-poppins-sb fs-18">Revenue (USD)</span>
      </div>
    </div>
  </div>
</template>

<script>
import SuInfoItem from "@/components/SuInfoItem";
export default {
  name: 'SuRevenueBox',
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
