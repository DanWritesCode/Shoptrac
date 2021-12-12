<template>
  <div class="card bg-box p-lg-4 p-3 h-100">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-18">Cost of Goods Sold</p>
      </div>
    </div>
    <div class="row align-items-center mt-2">
      <div class="col-5">
        <span class="ls-2 ff-poppins-sb">Expense</span>
      </div>
      <div class="d-none d-bp d-xl-block col-lg-4">
        <span class="ls-2 ff-poppins-sb">
          Amount
        </span>
      </div>
      <div class="col-lg-3 col-md-2 d-none d-lg-block">
        <span class="ls-2 ff-poppins-sb">% of Order</span>
      </div>
    </div>
    <div class="mb-4">
      <su-info-item name="test" colorClass="is-red" time="5123512">
      </su-info-item>
    </div>
  </div>
</template>

<script>
import SuInfoItem from "@/components/SuInfoItem";
export default {
  name: 'SuCogsBox',
  components: {SuInfoItem},
  props: {
    hasMore: {
      type: Boolean,
      default: false
    }
  }
}
</script>