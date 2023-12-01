connection "btp" {
  plugin = "ajmaradiaga/btp"

  # You will need to create a service key for the Cloud Management Service. You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis.

  # URL of the Accounts Service. Required.
  # This can also be set via the `BTP_CIS_ACCOUNTS_SERVICE_URL` environment variable.
  cis_accounts_service_url = "https://accounts-service.cfapps.[region].hana.ondemand.com"
  
  # URL of the Entitlements Service. Required.
  # This can also be set via the `BTP_CIS_ENTITLEMENTS_SERVICE_URL` environment variable.
  cis_entitlements_service_url = "https://entitlements-service.cfapps.[region].hana.ondemand.com"

  # Access token to communicate with the Cloud Management Service APIs. At the moment only access token is supported. You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis. Required.
  # This can also be set via the `BTP_CIS_ACCESS_TOKEN` environment variable.
  cis_access_token = "eyJhbGciOiDBNsO0JxFoAaodkDJ3Pmk7cFEsEr5ml5BwNWEafrEjy8Hsxt2mVACpD8B4AIPpRuMoGE71qXGoPcW0vCugceTwN4C3xM8qYmH7DLQrdVIlSX6kydYxnRNjSO8je56ckA4oTC8wm2E2clClPhinDBN6DxHhXlB0eVJnerl4ONpxaH43PYXmHjIsArTuBGK6nCFtApIGN1OvMDPmjHFtOjNgcPCPC5GDXTt5oaB6M2gUrfD5QQVGA7L6yFQlXYPvF6BSyMxpoQXywMUwYA6oqV"
 
}