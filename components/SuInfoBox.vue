<template>
  <div class="card bg-box mt-4 p-lg-4 p-3">
    <div class="row align-items-center mb-lg-3">
      <div class="col-4 col-lg-3">
        <span class="ls-2">Name</span>
      </div>
      <div class="d-none d-bp d-xl-block col-lg-4">
        <span class="ls-2">
          Date &amp; Time
        </span>
      </div>
      <div class="col-lg-3 col-md-2 d-none d-lg-block">
        <span class="ls-2">Status</span>
      </div>
      <div
        class="col-12 col-lg-2 text-lg-right text-left order-lg-last order-first"
      >
        <span class="ff-poppins-sb fs-18">History</span>
      </div>
    </div>
    <div class="content-box">
      <su-info-item name="test" colorClass="is-blue" time="5123512">
      </su-info-item>
    </div>
    <div v-if="hasMore" class="row align-content-end justify-content-end">
      <div class="col-12 text-right">
        <n-link
          to="example-path"
          class="ff-poppins text-uppercase text-white text-red-hover d-block mt-3 mr-3"
        >
          View More
        </n-link>
      </div>
    </div>
  </div>
</template>

<script>
import SuInfoItem from "@/components/SuInfoItem";
export default {
  name: 'SuInfoBox',
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
