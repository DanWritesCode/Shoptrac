<template>
  <div class="card bg-box p-lg-4 p-3 h-100">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-18">{{title}}</p>
      </div>
    </div>
    <div v-for="(key, val) in data">
      <div class="row align-items-center mb-lg-3 mt-2">
        <div class="col-6">
          <h2 class="text-white">{{key}}</h2>
        </div>
        <div class="col-6 text-right order-last">
          <span class="ff-poppins-sb fs-18">{{val}}</span>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import SuInfoItem from "@/components/SuInfoItem";
export default {
  name: 'SuMultilineInfoBox',
  components: {SuInfoItem},
  props: {
    hasMore: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: ""
    },
    data: {
      type: Array,
      default: {"meme": "sauce", "doge": "meme"}
    }
  }
}
</script>