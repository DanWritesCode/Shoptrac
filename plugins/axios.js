export default function({ $axios, redirect }) {
    $axios.interceptors.response.use(function (response) {
        return response;
    }, function (error) {
        if(error.response.status == 429) return $nuxt.$toastr('error', "Too many requests - try again in a few minutes", "Error");
        if (error.response.data.message == "You tried. You failed.") return Promise.reject(error);
        $nuxt.$toastr('error', error.response.data.message, "Error");
        if(error.response.data.message == "You need to agree to our terms of service before you can use our site") $nuxt.$router.push('/tos-gateway')
        if(error.response.status == 403 || error.response.status == 401) {
            localStorage.setItem('loggedIn', false)
            localStorage.setItem('username', '')
            localStorage.setItem('token', '')
            $nuxt.$router.push('/sign-in')
        }
        return Promise.reject(error);
      });

  $axios.interceptors.request.use((request) => requestHandler(request))

  const isHandlerEnabled = (config = {}) => {
    return config.hasOwnProperty('handlerEnabled') && !config.handlerEnabled
      ? false
      : true
  }

  const requestHandler = (request) => {
    if (isHandlerEnabled(request)) {
        request.headers['Authorization'] = localStorage.getItem("token"); 
        request.headers['cfg-ray'] = localStorage.getItem('cf-ray')
        request.headers['cfg-ray2'] = localStorage.getItem('cf-ray2')
    }
    return request
  }
}
