USE [ImportDataMobius]
GO

/****** Object:  Table [dbo].[CheckOGL]    Script Date: 19/06/2024 18:12:49 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[CheckOGL](
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
	[AmountEnteredDebit] [nvarchar](max) NULL,
	[AmountEnteredCredit] [nvarchar](max) NULL,
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


