connection "btp" {
  plugin = "ajmaradiaga/btp"

  # User email used to log in to SAP BTP. Required.
  # This can also be set via the `BTP_USERNAME` environment variable.
  # username = "user@domain.com"

  # User password used to log in to SAP BTP. Required.
  # This can also be set via the `BTP_PASSWORD` environment variable.
  # password = "My-BTP-Passw0rd"

  # You will need to create a service key for the Cloud Management Service. Required.
  # You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis.
  # This can also be set via the `BTP_CIS_SERVICE_KEY_PATH` environment variable. Required.
  # cis_service_key_path = "~/service_keys/cis_global.json"

  # URL of the Accounts Service. Optional.
  # This can also be set via the `BTP_CIS_ACCOUNTS_SERVICE_URL` environment variable.
  # cis_accounts_service_url = "https://accounts-service.cfapps.[region].hana.ondemand.com"
  
  # URL of the Entitlements Service. Optional.
  # This can also be set via the `BTP_CIS_ENTITLEMENTS_SERVICE_URL` environment variable.
  # cis_entitlements_service_url = "https://entitlements-service.cfapps.[region].hana.ondemand.com"

  # Access token to communicate with the Cloud Management Service APIs. Optional.
  # You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis. 
  # If no access token is provided, the plugin will try getting an access token using the details provided in cis_client_id, cis_client_secret, cis_token_url, username, password.
  # This can also be set via the `BTP_CIS_ACCESS_TOKEN` environment variable.
  # cis_access_token = "eyJhbGciOiDBNsO0JxFoAaodkDJ3Pmk7cFEsEr5ml5BwNWEafrEjy8Hsxt2mVACpD8B4AIPpRuMoGE71qXGoPcW0vCugceTwN4C3xM8qYmH7DLQrdVIlSX6kydYxnRNjSO8je56ckA4oTC8wm2E2clClPhinDBN6DxHhXlB0eVJnerl4ONpxaH43PYXmHjIsArTuBGK6nCFtApIGN1OvMDPmjHFtOjNgcPCPC5GDXTt5oaB6M2gUrfD5QQVGA7L6yFQlXYPvF6BSyMxpoQXywMUwYA6oqV"

  # A service key (https://help.sap.com/docs/btp/sap-business-technology-platform/creating-service-keys) for the Central Management Service is required to get an access token.  
  # The values for cis_client_id, cis_client_secret, cis_token_url can be retrieved from it. 

  # OAuth 2.0 Client ID (field in the service key: uua.clientid). Optional.
  # This can also be set via the `BTP_CIS_CLIENT_ID` environment variable.
  # cis_client_id = "sb-ut-6y0m0wr1-ai10-1a56-4bv1-52m478h011g6-clone!b017880|cis-central!b14"

  # OAuth 2.0 Client Secret (field in the service key: uua.clientsecret). Optional.
  # This can also be set via the `BTP_CIS_CLIENT_SECRET` environment variable. Optional.
  # cis_client_secret = "44s32264-285r-7101-g2n0-g036p1m214us$vm2_UCHAN_mX1FOW0lklBj2-igBsOUG77G-nE1TWsEu="

  # OAuth 2.0 Token URL (field in the service key: uua.url). Optional.
  # The value in the service key doesn't contain the path, /oauth/token, if not specified the plugin will append it automatically.
  # This can also be set via the `BTP_CIS_TOKEN_URL` environment variable. Optional.
  # cis_token_url = "https://[global-account-subdomain].authentication.[region].hana.ondemand.com"
<<<<<<< HEAD
=======

  # User Email used to log in to SAP BTP.
  # This can also be set via the `BTP_USERNAME` environment variable. Optional.
  # username = "user@domain.com"

  # User Password used to log in to SAP BTP.
  # This can also be set via the `BTP_PASSWORD` environment variable. Optional.
  # password = "My-BTP-Passw0rd"
>>>>>>> 2f53a7e (Including additional config supported)
}