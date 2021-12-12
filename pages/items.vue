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
    </div>
    <div class="row mw-100 mr-0">
      <div class="col-12 mr-0 pr-0">
        <su-dashboard-chart></su-dashboard-chart>
      </div>
    </div>
    <div class="row">
      <!-- Double mt-4 because the above boxes have weird sizing issues. Oh well... -->
      <div class="col-12 mt-4"><h3>Item Sales Information</h3></div>
    </div>
    <div class="row">
      <div class="col-12 col-lg-6">
        <su-multiline-info-box title="Average Item Price" :data="aip">
        </su-multiline-info-box>
      </div>
      <div class="col-12 col-lg-6">
        <su-multiline-info-box title="Inventory Value" :data="inv">
        </su-multiline-info-box>
      </div>
    </div>
    <div class="row mt-4">
      <div class="col-12 col-lg-6">
        <su-list-box title="Top Selling Items" :data="aip">
        </su-list-box>
      </div>
      <div class="col-12 col-lg-6">
        <su-list-box title="Top Selling Collections" :data="inv">
        </su-list-box>
      </div>
      <!-- Top Selling Items -->
      <!-- Top Selling Collections -->
    </div>
  </div>
</template>

<script>
import Alert from '../components/Alert'
import SuDashboardChart from "@/components/SuDashboardChart";
import SuMultilineInfoBox from "@/components/SuMultilineInfoBox";
import SuListBox from "@/components/SuListBox";

export default {
  name: 'Items',
  components: {SuListBox, SuMultilineInfoBox, SuDashboardChart, Alert },
  layout: 'dashboard',
  transition: 'fade',
  head() {
    return {
      title: 'StonksUp | Items'
    }
  },
  data() {
    return {
      notices: [],
      aip: {"Average Item Price": "$55.55"},
      inv: {"Inventory Value": "$512551.55"}
    }
  },
  methods: {

  }
}
</script>