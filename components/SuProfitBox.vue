<template>
  <div class="card bg-box p-4 p-3 mb-5 mb-md-0">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-18">Profit/Loss Breakdown</p>
      </div>
    </div>
    <div class="row align-items-center mb-3 mt-2">
      <div class="col-6">
        <h1 class="text-green">$1234.56</h1>
      </div>
      <div class="col-6 text-right order-last order-first">
        <span class="ff-poppins-sb fs-18">Profit (USD)</span>
      </div>
    </div>

    <div class="row align-items-center mb-1">
      <div class="col-6">
        <h1 class="text-green">22.8%</h1>
      </div>
      <div class="col-6 text-right order-last order-first">
        <span class="ff-poppins-sb fs-18">Profit Margin (%)</span>
      </div>
    </div>
  </div>
</template>

<script>
import SuInfoItem from "@/components/SuInfoItem";
export default {
  name: 'SuProfitBox',
  components: {SuInfoItem},
  props: {
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
