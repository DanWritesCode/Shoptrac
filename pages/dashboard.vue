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
import {API_URL} from "@/config";

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
  async mounted() {
    await this.$axios
        .$get(`${API_URL}/summary`, {
          headers: {
            Authorization: localStorage.getItem('token')
          }
        })
        .then((res) => {
          this.populateDashboardSummary(res);
        })
        .catch((res) => {
        })
  },
  data() {
    return {
      summaryData: {},
      notices: [],
      orderBreakdown: {},
      revenueBreakdown: {},
      profitBreakdown: {},

      cogsCostCol: ["Expense", "Amount", "% Revenue"],
      marketingCostCol: ["Platform", "Spend", "% Revenue"],
      recurringCostCol: ["Expense", "Amount", "% Revenue"],

      cogsCost: [],
      marketingCost: [],
      recurringCost: []
    }
  },
  methods: {
    formatNumber(number) {
      return Number(number).toLocaleString(undefined, {minimumFractionDigits: 2, maximumFractionDigits: 2});
    },
    populateDashboardSummary(summaryData) {
      this.orderBreakdown = {"Orders": summaryData.orders, "Average Order Value": "$" + this.formatNumber(summaryData.aov)}
      this.revenueBreakdown = {"Revenue": "$" + this.formatNumber(summaryData.revenue), "Expenses": "$" + this.formatNumber(summaryData.expenses)}
      this.profitBreakdown = {"Profit (USD)": "$" + this.formatNumber(summaryData.profit), "Profit Margin": this.formatNumber(summaryData.profitMargin) + "%"}

      this.cogsCost = summaryData.groupedExpenses.cogs;
      for(let i = 0; i < this.cogsCost.length; i++) {
        this.cogsCost[i].amount = "$" + this.formatNumber(this.cogsCost[i].amount);
        this.cogsCost[i].percentage = this.formatNumber(this.cogsCost[i].percentage) + "%";
      }
      this.marketingCost = summaryData.groupedExpenses.marketing;
      for(let i = 0; i < this.marketingCost.length; i++) {
        this.marketingCost[i].amount = "$" + this.formatNumber(this.marketingCost[i].amount);
        this.marketingCost[i].percentage = this.formatNumber(this.marketingCost[i].percentage) + "%";
      }
      this.recurringCost = summaryData.groupedExpenses.recurring;
      for(let i = 0; i < this.recurringCost.length; i++) {
        this.recurringCost[i].amount = "$" + this.formatNumber(this.recurringCost[i].amount);
        this.recurringCost[i].percentage = this.formatNumber(this.recurringCost[i].percentage) + "%";
      }

    }
  }
}
</script>

<style scoped></style>
