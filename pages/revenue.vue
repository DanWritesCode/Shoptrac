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
      <div class="col-12 mt-4"><h3>Revenue</h3></div>
    </div>
    <div class="row">
      <div class="col-12 col-lg-10">
        <su-list-box title="Store Revenue" :columns="revenueColumns" :data="revenueData" color-class="is-green"></su-list-box>
      </div>
    </div>
  </div>
</template>

<script>
import Alert from '../components/Alert'
import SuDashboardChart from "@/components/SuDashboardChart";
import SuListBox from "@/components/SuListBox";
import {API_URL} from "@/config";

export default {
  name: 'Revenue',
  components: {SuListBox, SuDashboardChart, Alert },
  layout: 'dashboard',
  transition: 'fade',
  head() {
    return {
      title: 'StonksUp | All Revenue'
    }
  },
  data() {
    return {
      notices: [],
      revenueData: [{}],
      revenueColumns: ["Item Sales", "Shipping Charged", "Taxes Collected", "Tips"]
    }
  },
  async mounted() {
    await this.$axios
        .$get(`${API_URL}/revenue`, {
          headers: {
            Authorization: localStorage.getItem('token')
          }
        })
        .then((res) => {
          res.sales = "$" + res.sales.toFixed(2).toLocaleString();
          res.shippingCharged = "$" + res.shippingCharged.toFixed(2).toLocaleString();
          res.taxesCollected = "$" + res.taxesCollected.toFixed(2).toLocaleString();
          res.tips = "$" + res.tips.toFixed(2).toLocaleString();

          this.revenueData = [res];
        })
        .catch((res) => {
        })
  },
  methods: {

  }
}
</script>