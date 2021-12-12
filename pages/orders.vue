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
      <div class="col-12 mt-4"><h3>Order Information</h3></div>
    </div>
    <div class="row">
      <!-- Orders -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Orders" :data="aip">
        </su-multiline-info-box>
      </div>

      <!-- Average Order Value -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Average Order Value" :data="aip">
        </su-multiline-info-box>
      </div>

      <!-- Order Margin -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Order Margin" :data="aip">
        </su-multiline-info-box>
      </div>

      <!-- Refunds -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Refunds" :data="aip">
        </su-multiline-info-box>
      </div>
    </div>
    <!-- Recent Orders (Amount, Shipping, COGS, Profit/Margin) -->
    <div class="row mt-4">
      <div class="col-12 col-lg-12">
        <su-list-box title="Order List" :data="inv">
        </su-list-box>
      </div>
    </div>

  </div>
</template>

<script>
import Alert from '../components/Alert'
import SuDashboardChart from "@/components/SuDashboardChart";
import SuMultilineInfoBox from "@/components/SuMultilineInfoBox";
import SuListBox from "@/components/SuListBox";

export default {
  name: 'Orders',
  components: { SuDashboardChart, SuMultilineInfoBox, SuListBox, Alert },
  layout: 'dashboard',
  transition: 'fade',
  head() {
    return {
      title: 'StonksUp | Orders'
    }
  },
  data() {
    return {
      notices: [],
      inv: [{}]
    }
  },
  methods: {

  }
}
</script>