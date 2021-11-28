<template>
  <div>
    <div class="row">
      <div class="col-12">
        <alert :type="'alert-info'">Test Banner/Alert</alert>
        <div v-for="item in notices">
          <alert :type="'alert-info'" v-html="item.message">
          </alert>
        </div>
      </div>
      <div class="col-12 col-lg-6">
        <su-info-box>

        </su-info-box>
      </div>
      <div class="col-12 col-lg-6">
        <su-info-box>

        </su-info-box>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
      </div>
    </div>
  </div>
</template>

<script>
import Alert from '../components/Alert'
import SuInfoBox from "@/components/SuInfoBox";

export default {
  name: 'Dashboard',
  components: {SuInfoBox, Alert },
  layout: 'dashboard',
  transition: 'fade',
  head() {
    return {
      title: 'StonksUp | Dashboard'
    }
  },
  data() {
    return {
    }
  },
  methods: {

  }
}
</script>

<style scoped></style>
