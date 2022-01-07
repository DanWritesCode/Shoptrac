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
        <su-multiline-info-box title="Orders" :data="orders">
        </su-multiline-info-box>
      </div>

      <!-- Average Order Value -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Average Order Value" :data="aov">
        </su-multiline-info-box>
      </div>

      <!-- Order Margin -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Profit Margin" :data="orderMargin">
        </su-multiline-info-box>
      </div>

      <!-- Refunds -->
      <div class="col-12 col-lg-3">
        <su-multiline-info-box title="Refunds" :data="refunds">
        </su-multiline-info-box>
      </div>
    </div>
    <!-- Recent Orders (Amount, Shipping, COGS, Profit/Margin) -->
    <div class="row mt-4">
      <div class="col-12 col-lg-12">
        <su-list-box title="Order List" :columns="orderListCol" :data="orderList">
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
      orders: {},
      aov: {},
      orderMargin: {},
      refunds: {},

      orderList: [{}],
      orderListCol: ["Order ID", "Date", "Item Count", "Country", "Total Amount", "COGS"],

      notices: [],
    }
  },
  async mounted() {
    await this.$axios
        .$get(`${API_URL}/orders`, {
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
  methods: {
    formatNumber(number) {
      return Number(number).toLocaleString(undefined, {minimumFractionDigits: 2, maximumFractionDigits: 2});
    },
    populate(data) {
      this.orders = {"Orders": data.orders};
      this.aov = {"AOV": "$" + this.formatNumber(data.aov)};
      this.orderMargin = {"Margin": this.formatNumber(data.margin) + "%"};
      this.refunds = {"Refunded": "$" + this.formatNumber(data.refunds)};

      this.orderList = data.orderList;
      for(let i = 0; i < this.orderList.length; i++) {
        this.orderList[i].orderId = "#" + this.orderList[i].orderId;
        this.orderList[i].amount = "$" + this.formatNumber(this.orderList[i].amount);
        this.orderList[i].cogs = "$" + this.formatNumber(this.orderList[i].cogs);
        this.orderList[i].date = new Date(this.orderList[i].date*1000).toLocaleString();
        delete this.orderList[i].paymentGateway;
        delete this.orderList[i].discount;
        delete this.orderList[i].shippingCharged;
        delete this.orderList[i].subtotal;
        delete this.orderList[i].taxesCharged;
      }
    },
  }
}
</script>