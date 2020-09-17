module.exports = {
  publicPath: process.env.BASE_URL,

  // So we don't get an 'Invalid Host' error message when we
  // try to run this in production mode. Notes here: https://bit.ly/3lI7sig
  // don't set the scheme (duh) because it's a host, not a fqdn
  devServer: {
    public: 'skillcalc5e.cc'
  }
}
