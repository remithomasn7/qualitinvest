from alphavantage_api_client import AlphavantageClient, GlobalQuote, Quote, AccountingReport, CompanyOverview, Ticker
import logging
from decimal import Decimal

def sample_ticker_usage():
    """
    combine all financial statements (income, cash flow, earnings and balance sheet) for both
     quarterly and annual reports. There migth be times you want to store/visualize them together
     thus providing another dimensionality of the data
    Returns:

    """
    aapl = (Ticker()
            .create_client()
            .should_retry_once()  # auto retry when limit reached
            .from_symbol("AAPL")  # define the company of interest
            .fetch_accounting_reports()  # make the call to alpha vantage api
            .correlate_accounting_reports()  # combines all 4 financial statements
            )

    # get the individual financial statements if needed
    earnings = aapl.get_earnings()
    income_statement = aapl.get_income_statement()
    balance_sheet = aapl.get_balance_sheet()
    cash_flow = aapl.get_cash_flow()

    # get the combined financial statements and iterate easy
    correlated_financial_statements = aapl.get_correlated_reports() # both quarterly and annually

    for fiscal_date_ending in correlated_financial_statements:
        combined_financial_statements = correlated_financial_statements[fiscal_date_ending]
        print(f"{fiscal_date_ending} = {combined_financial_statements}")
    

if __name__ == "__main__":
    sample_ticker_usage()
    import pdb; pdb.set_trace()