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
      <div class="col-12 mt-4"><h3>Customer Information</h3></div>
    </div>
    <div class="row">
      <!-- Total Customers -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Total Customers" :data="totalCustomers">
        </su-multiline-info-box>
      </div>

      <!-- New Customers -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="New Customers" :data="newCustomers">
        </su-multiline-info-box>
      </div>

      <!-- Returning Customers -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Returning Customers" :data="returningCustomers">
        </su-multiline-info-box>
      </div>

      <!-- Customer Spending Range -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Customer Spending Range" :data="customerSpendRange">
        </su-multiline-info-box>
      </div>
    </div>
    <!-- Customer List (Name, Orders, Amount Spent, Returning or New) -->
    <div class="row mt-4">
      <div class="col-12 col-lg-12">
        <su-list-box title="Customer List" :data="inv">
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
  name: 'Customers',
  components: {SuListBox, SuMultilineInfoBox, SuDashboardChart, Alert },
  layout: 'dashboard',
  transition: 'fade',
  head() {
    return {
      title: 'StonksUp | Customers'
    }
  },
  data() {
    return {
      notices: [],
      totalCustomers: {"Total Customers": 0},
      newCustomers: {"New Customers": 0},
      returningCustomers: {"Returning Customers": 0},
      customerSpendRange: {"Highest Spender": 0, "Lowest Spender": 0},
    }
  },
  methods: {

  }
}
</script>