USE [ImportDataMobius]
GO

/****** Object:  Table [dbo].[ReportOGL]    Script Date: 19/06/2024 8:31:08 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[ReportOGL](
	[Company] [nvarchar](max) NULL,
	[RC] [nvarchar](max) NULL,
	[OC] [nvarchar](max) NULL,
	[Channel] [nvarchar](max) NULL,
	[ProductCode] [nvarchar](max) NULL,
	[DrCr] [nvarchar](max) NULL,
	[GLAccountNo] [nvarchar](max) NULL,
	[GLAccountName] [nvarchar](max) NULL,
	[Activity] [nvarchar](max) NULL,
	[Tax] [nvarchar](max) NULL,
	[InterCo] [nvarchar](max) NULL,
	[Future1] [nvarchar](max) NULL,
	[Future2] [nvarchar](max) NULL,
	[Currency] [nvarchar](max) NULL,
	[AmountEnteredDebit] [decimal](18, 2) NULL,
	[AmountEnteredCredit] [decimal](18, 2) NULL,
	[LoanAccountNumber] [nvarchar](max) NULL,
	[GroupReferenceNumber] [nvarchar](max) NULL,
	[OriginalTransactionReferenceNumber] [nvarchar](max) NULL,
	[TransactionReferenceNumber] [nvarchar](max) NULL,
	[GLGroupCodeCoA] [nvarchar](max) NULL,
	[TransactionPostingDate] [nvarchar](max) NULL,
	[EffectiveDate] [nvarchar](max) NULL,
	[JournalEntryDescription] [nvarchar](max) NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO


