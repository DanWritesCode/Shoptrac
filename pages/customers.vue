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
        <su-multiline-info-box title="Total Customers" :data="totalCustomers" class="h-100">
        </su-multiline-info-box>
      </div>

      <!-- New & Returning Customers -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="New vs Returning Customers" :data="newReturningCustomers">
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
      <div class="col-12 col-lg-9">
        <su-list-box title="Most Valuable Customers" :columns="customersCol" :data="customers">
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
import {API_URL} from "@/config";

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
  async mounted() {
    await this.$axios
        .$get(`${API_URL}/customers`, {
          headers: {
            Authorization: localStorage.getItem('token')
          }
        })
        .then((res) => {
          this.data = res;
          this.populate(res);
        })
        .catch((res) => {
        })
  },
  data() {
    return {
      notices: [],
      totalCustomers: {"Total Customers": 0},
      newReturningCustomers: {"New Customers": 0, "Returning Customers": 0},
      customerSpendRange: {"Highest Spender": 0, "Lowest Spender": 0},
      customersCol: ["Customer Name", "Country", "Orders Placed", "Total Amount Spent"],
      customers: [{}],
    }
  },
  methods: {
    populate(data) {
      this.totalCustomers = {"Total Customers": data.totalCustomers};
      this.newReturningCustomers = {"New Customers": data.newCustomers, "Returning Customers": data.returningCustomers};
      this.customerSpendRange = {"Highest Spender": "$" + data.highestSpender, "Lowest Spender": "$" + data.lowestSpender};

      this.customers = data.topCustomerList;
      for(let i = 0; i < this.customers.length; i++) {
        this.customers[i].amountSpent = "$" + this.customers[i].amountSpent.toFixed(2).toLocaleString();

      }
    }
  }
}
</script>