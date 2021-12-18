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
    <div class="row mt-4">
      <!-- Double mt-4 because the above boxes have weird sizing issues. Oh well... -->
      <div class="col-12 mt-4"><h3>Income</h3></div>
    </div>
    <div class="row">

      <div class="col-12 col-lg-4">
        <su-dashboard-box title="Profit Breakdown" :data="profitBreakdown" color="text-green">

        </su-dashboard-box>
      </div>
      <div class="col-12 col-lg-4">
        <su-dashboard-box title="Revenue Breakdown" :data="revenueBreakdown">

        </su-dashboard-box>
      </div>
      <div class="col-12 col-lg-4">
        <su-dashboard-box title="Order Breakdown" :data="orderBreakdown">
        </su-dashboard-box>
      </div>
    </div>
    <div class="row">
      <div class="col-12 mt-4"><h3>Expenses</h3></div>
    </div>
    <div class="row">
      <div class="col-12 col-lg-4">
        <su-dashboard-expense-box color="text-red" title="Cost of Goods Sold" :columns="cogsCostCol" :data="cogsCost">

        </su-dashboard-expense-box>
      </div>
      <div class="col-12 col-lg-4">
        <su-dashboard-expense-box color="text-red" title="Marketing Costs" :columns="marketingCostCol" :data="marketingCost">

        </su-dashboard-expense-box>
      </div>
      <div class="col-12 col-lg-4">
        <su-dashboard-expense-box  color="text-red" title="Recurring Costs" :columns="recurringCostCol" :data="recurringCost">

        </su-dashboard-expense-box >
      </div>
    </div>
  </div>
</template>

<script>
import Alert from '../components/Alert'
import SuDashboardChart from "@/components/SuDashboardChart";
import SuDashboardBox from "@/components/SuMultilineInfoBox";
import SuDashboardExpenseBox from "@/components/SuDashboardExpenseBox";

export default {
  name: 'Dashboard',
  components: {
    SuDashboardExpenseBox,
    SuDashboardBox,
    SuDashboardChart, Alert },
  layout: 'dashboard',
  transition: 'fade',
  head() {
    return {
      title: 'StonksUp | Dashboard'
    }
  },
  data() {
    return {
      notices: [],
      orderBreakdown: {"Orders": 524, "Average Order Value": "$112.43"},
      revenueBreakdown: {"Revenue": "$51,225.25", "Expenses": "$25,512.12"},
      profitBreakdown: {"Profit (USD)": "$1,234.56", "Profit Margin": "55.5%"},

      cogsCostCol: ["Expense", "Amount", "% Revenue"],
      marketingCostCol: ["Platform", "Spend", "% Revenue"],
      recurringCostCol: ["Expense", "Amount", "% Revenue"],

      cogsCost: [{0: "Items Sold", 1: "$4444.44", 2: "22.8%"}],
      marketingCost: [{0: "Facebook", 1: "$2222.22", 2: "11.4%"}],
      recurringCost: [{0: "Shopify Plan", 1: "$29.99", 2: "0.8%"}]
    }
  },
  methods: {

  }
}
</script>

<style scoped></style>
