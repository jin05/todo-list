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
        },
      };
    default:
      return {
        // environment varibales for production
        env: {
          IDP_DOMAIN: "local-todo-list.auth.ap-northeast-1.amazoncognito.com",
          USER_POOL_ID: "ap-northeast-1_axGAoF8m8",
          USER_POOL_CLIENT_ID: "6u3mjd9gmvo638193q9ci4lmi6",
          REDIRECT_SIGN_IN: "http://localhost:3000/token",
          REDIRECT_SIGN_OUT: "http://localhost:3000/",
          AUTH_COOKIE_DOMAIN: "localhost",
        },
      };
  }
};
