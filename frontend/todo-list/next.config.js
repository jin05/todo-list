const { PHASE_DEVELOPMENT_SERVER } = require("next/constants");

module.exports = (phase) => {
  switch (phase) {
    case PHASE_DEVELOPMENT_SERVER:
      return {
        // environment varibales for local development
        env: {
          IDP_DOMAIN: "local-todo-list.auth.ap-northeast-1.amazoncognito.com",
          USER_POOL_ID: "ap-northeast-1_axGAoF8m8",
          USER_POOL_CLIENT_ID: "6u3mjd9gmvo638193q9ci4lmi6",
          REDIRECT_SIGN_IN: "http://localhost:3000/token",
          REDIRECT_SIGN_OUT: "http://localhost:3000/",
          AUTH_COOKIE_DOMAIN: "localhost",
          API_HOST: "http://localhost:8080",
        },
      };
    default:
      return {
        // environment varibales for production
        env: {
          IDP_DOMAIN: "prod-todo-list.auth.ap-northeast-1.amazoncognito.com",
          USER_POOL_ID: "ap-northeast-1_3yXLWYIMZ",
          USER_POOL_CLIENT_ID: "79ruu0siuha5neotiggbfvaj3l",
          REDIRECT_SIGN_IN: "https://d2semoivot4v0t.cloudfront.net/token",
          REDIRECT_SIGN_OUT: "https://d2semoivot4v0t.cloudfront.net/",
          AUTH_COOKIE_DOMAIN: "d2semoivot4v0t.cloudfront.net",
          API_HOST:
            "https://xn8hqr060e.execute-api.ap-northeast-1.amazonaws.com",
        },
      };
  }
};
