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
      <div class="col-12 mt-4"><h3>Expenses</h3></div>
    </div>
    <div class="row">
      <div class="col-12">
        <su-list-box title="Store Expenses Summary" :columns="expensesColumns" :data="expensesData" color-class="is-red"></su-list-box>
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
  name: 'Expenses',
  components: {SuListBox, SuDashboardChart, Alert },
  layout: 'dashboard',
  transition: 'fade',
  head() {
    return {
      title: 'StonksUp | All Expenses'
    }
  },
  data() {
    return {
      notices: [],
      expensesData: [{}],
      expensesColumns: ["Payment Processing", "COGS (Product)", "COGS (Shipping)", "Taxes Forwarded", "Marketing", "Recurring Costs", "One-time Costs"]
    }
  },
  async mounted() {
    await this.$axios
        .$get(`${API_URL}/expenses`, {
          headers: {
            Authorization: localStorage.getItem('token')
          }
        })
        .then((res) => {
          for(let k in res)
            if(typeof res[k] !== 'object')
              res[k] = "$" + this.formatNumber(res[k])

          let totalMarketing = 0;
          let totalRecurring = 0;
          for(let k in res.marketing)
            totalMarketing += res.marketing[k].amount;

          for(let k in res.recurring)
            totalRecurring += res.recurring[k].amount;

          delete res.marketing;
          res.marketing = "$" + this.formatNumber(totalMarketing);

          delete res.recurring;
          res.recurring = "$" + this.formatNumber(totalRecurring);

          res.oneTime = "$0.00";

          this.expensesData = [res];
        })
        .catch((res) => {
        })
  },
  methods: {
    formatNumber(number) {
      return Number(number).toLocaleString(undefined, {minimumFractionDigits: 2, maximumFractionDigits: 2});
    },
  }
}
</script>