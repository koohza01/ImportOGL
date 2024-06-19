select TransactionReferenceNumber,Sum(AmountEnteredCredit) as credit,Sum(AmountEnteredDebit) as debit 
from ReportOGL group by TransactionReferenceNumber
having Sum(AmountEnteredCredit) <> Sum(AmountEnteredDebit)


select SUM(amtcredit) as amtcredit,SUM(amtdebit) as amtdebit from(select TransactionReferenceNumber,Sum(AmountEnteredCredit) as amtcredit,Sum(AmountEnteredDebit) as amtdebit
from ReportOGL group by TransactionReferenceNumber
having Sum(AmountEnteredCredit) <> Sum(AmountEnteredDebit)) sss

select Sum(AmountEnteredDebit) as amtdebit,Sum(AmountEnteredCredit) as amtcredit,  Sum(AmountEnteredDebit) - Sum(AmountEnteredCredit) as diff
from ReportOGL 

