<template>
  <div class="card bg-box p-lg-4 p-3 h-100">
    <div class="row text-center">
      <div class="mx-auto">
        <p class="ff-poppins-sb fs-20">{{title}}</p>
      </div>
    </div>
    <div class="row align-items-center text-center mt-2">
      <div class="col" v-for="column in columns">
        <span class="ls-2 ff-poppins-sb fs-18">{{column}}</span>
      </div>
    </div>
    <div class="mb-4">
      <su-info-item :color="color" :data="item" v-for="item in data">
      </su-info-item>
    </div>
  </div>
</template>

<script>
import SuInfoItem from "@/components/SuInfoItem";
export default {
  name: 'SuDashboardExpenseBox',
  components: {SuInfoItem},
  props: {
    title: {
      type: String,
      default: ""
    },
    columns: {
      type: Array,
      default: []
    },
    data: {
      type: Array,
      default: []
    },
    color: {
      type: String,
      default: "is-white"
    }
  }
}
</script>