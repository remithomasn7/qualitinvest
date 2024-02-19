from alphavantage_api_client import AlphavantageClient, AccountingReport
import logging
from decimal import Decimal

def calculate_free_cash_flow_last_fiscal_year(symbol):
    client = AlphavantageClient()
    one_billion = Decimal(1000000000.00) # just make the numbers smaller and easier to read :-)
    cash_flow_statements = client.get_cash_flow(symbol)
    last_year_cash_flow_statement = cash_flow_statements.get_most_recent_annual_report()
    last_year_fcf = calc_free_cash_flow(last_year_cash_flow_statement, one_billion)
    fiscalDateEnding = last_year_cash_flow_statement.get('fiscalDateEnding', 'Not Provided')
    print(f"FCF as of {fiscalDateEnding} for {symbol} was ${last_year_fcf}b")


def calc_free_cash_flow(free_cash_flow_statement: dict, multiple: Decimal) -> Decimal:
    capitalExpenditures = free_cash_flow_statement.get('capitalExpenditures', Decimal('0.00'))
    capitalExpenditures = Decimal(capitalExpenditures) / multiple

    operatingCashflow = free_cash_flow_statement.get('operatingCashflow', Decimal('0.00'))
    operatingCashflow = Decimal(operatingCashflow) / multiple

    free_cash_flow = operatingCashflow - capitalExpenditures
    print(f"operatingCashflow = ${operatingCashflow}b, capitalExpenditures = ${capitalExpenditures}b ")


    return free_cash_flow

def calc_free_cash_flow_per_share(symbol):
    client = AlphavantageClient()
    one_billion = Decimal(1000000000.00)
    cash_flow_statements = client.get_cash_flow(symbol)
    last_year_cash_flow_statement = cash_flow_statements.get_most_recent_annual_report()
    last_year_fcf = calc_free_cash_flow(last_year_cash_flow_statement, one_billion)
    fiscalDateEnding = last_year_cash_flow_statement.get('fiscalDateEnding', 'Not Provided')
    print(f"FCF as of {fiscalDateEnding} for {symbol} was ${last_year_fcf}b")

    # get shares out standing
    company_overview = client.get_company_overview(symbol)
    shares_outstanding = Decimal(company_overview.shares_outstanding) / one_billion
    fcf_per_share = round(last_year_fcf / shares_outstanding, 2)
    print(f"Free Cash Flow per share for {symbol} was ${fcf_per_share} having shares outstanding of {shares_outstanding}b")

if __name__ == "__main__":
    symbol = "AAPL"
    calc_free_cash_flow_per_share(symbol)
