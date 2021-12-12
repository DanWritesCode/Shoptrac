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
        <su-profit-box>

        </su-profit-box>
      </div>
      <div class="col-12 col-lg-4">
        <su-revenue-box>

        </su-revenue-box>
      </div>
      <div class="col-12 col-lg-4">
        <su-order-box>

        </su-order-box>
      </div>
    </div>
    <div class="row">
      <div class="col-12 mt-4"><h3>Expenses</h3></div>
    </div>
    <div class="row">
      <div class="col-12 col-lg-4">
        <su-cogs-box>

        </su-cogs-box>
      </div>
      <div class="col-12 col-lg-4">
        <su-marketing-expenses-box>

        </su-marketing-expenses-box>
      </div>
      <div class="col-12 col-lg-4">
        <su-recurring-expenses-box>

        </su-recurring-expenses-box>
      </div>
    </div>
  </div>
</template>

<script>
import Alert from '../components/Alert'
import SuInfoBox from "@/components/SuInfoBox";
import SuRevenueBox from "@/components/SuRevenueBox";
import SuProfitBox from "@/components/SuProfitBox";
import SuOrderBox from "@/components/SuOrderBox";
import SuCogsBox from "@/components/SuCogsBox";
import SuMarketingExpensesBox from "@/components/SuMarketingExpensesBox";
import SuRecurringExpensesBox from "@/components/SuRecurringExpensesBox";
import SuDashboardChart from "@/components/SuDashboardChart";

export default {
  name: 'Dashboard',
  components: {
    SuDashboardChart,
    SuRecurringExpensesBox,
    SuMarketingExpensesBox, SuCogsBox, SuOrderBox, SuProfitBox, SuRevenueBox, SuInfoBox, Alert },
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
    }
  },
  methods: {

  }
}
</script>

<style scoped></style>
