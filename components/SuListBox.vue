<template>
  <div class="card bg-box p-lg-4 p-3 h-100">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-18">{{title}}</p>
      </div>
    </div>

    <div class="row align-items-center mt-2">
      <div class="col text-center" v-for="colName in columns">
        <span class="ls-2 ff-poppins-sb">{{ colName }}</span>
      </div>
    </div>
    <div class="mb-4" v-for="(val) in data">
      <su-info-item :data="val" :colorClass="colorClass">
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
      default: []
    },
    colorClass: {
      type: String,
      default: "is-blue"
    },
    columns: {
      type: Array,
      default: []
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
