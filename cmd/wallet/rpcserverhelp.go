package wallet

// AUTOGENERATED by internal/rpchelp/genrpcserverhelp.go; do not edit.

func HelpDescsEnUS() map[string]string {
	return map[string]string{
		"addmultisigaddress":      "addmultisigaddress nrequired [\"key\",...] (\"account\")\n\nGenerates and imports a multisig address and redeeming script to the 'imported' account.\n\nArguments:\n1. nrequired (numeric, required)         The number of signatures required to redeem outputs paid to this address\n2. keys      (array of string, required) Pubkeys and/or pay-to-pubkey-hash addresses to partially control the multisig address\n3. account   (string, optional)          DEPRECATED -- Unused (all imported addresses belong to the imported account)\n\nResult:\n\"value\" (string) The imported pay-to-script-hash address\n",
		"createmultisig":          "createmultisig nrequired [\"key\",...]\n\nGenerate a multisig address and redeem script.\n\nArguments:\n1. nrequired (numeric, required)         The number of signatures required to redeem outputs paid to this address\n2. keys      (array of string, required) Pubkeys and/or pay-to-pubkey-hash addresses to partially control the multisig address\n\nResult:\n{\n \"address\": \"value\",      (string) The generated pay-to-script-hash address\n \"redeemScript\": \"value\", (string) The script required to redeem outputs paid to the multisig address\n}                         \n",
		"dumpprivkey":             "dumpprivkey \"address\"\n\nReturns the private key in WIF encoding that controls some wallet address.\n\nArguments:\n1. address (string, required) The address to return a private key for\n\nResult:\n\"value\" (string) The WIF-encoded private key\n",
		"getaccount":              "getaccount \"address\"\n\nDEPRECATED -- Lookup the account name that some wallet address belongs to.\n\nArguments:\n1. address (string, required) The address to query the account for\n\nResult:\n\"value\" (string) The name of the account that 'address' belongs to\n",
		"getaccountaddress":       "getaccountaddress \"account\"\n\nDEPRECATED -- Returns the most recent external payment address for an account that has not been seen publicly.\nA new address is generated for the account if the most recently generated address has been seen on the blockchain or in mempool.\n\nArguments:\n1. account (string, required) The account of the returned address\n\nResult:\n\"value\" (string) The unused address for 'account'\n",
		"getaddressesbyaccount":   "getaddressesbyaccount \"account\"\n\nDEPRECATED -- Returns all addresses strings controlled by a single account.\n\nArguments:\n1. account (string, required) Account name to fetch addresses for\n\nResult:\n[\"value\",...] (array of string) All addresses controlled by 'account'\n",
		"getbalance":              "getbalance (\"account\" minconf=1)\n\nCalculates and returns the balance of one or all accounts.\n\nArguments:\n1. account (string, optional)             DEPRECATED -- The account name to query the balance for, or \"*\" to consider all accounts (default=\"*\")\n2. minconf (numeric, optional, default=1) Minimum number of block confirmations required before an unspent output's value is included in the balance\n\nResult (account != \"*\"):\nn.nnn (numeric) The balance of 'account' valued in bitcoin\n\nResult (account = \"*\"):\nn.nnn (numeric) The balance of all accounts valued in bitcoin\n",
		"getbestblockhash":        "getbestblockhash\n\nReturns the hash of the newest block in the best chain that wallet has finished syncing with.\n\nArguments:\nNone\n\nResult:\n\"value\" (string) The hash of the most recent synced-to block\n",
		"getblockcount":           "getblockcount\n\nReturns the blockchain height of the newest block in the best chain that wallet has finished syncing with.\n\nArguments:\nNone\n\nResult:\nn.nnn (numeric) The blockchain height of the most recent synced-to block\n",
		"getinfo":                 "getinfo\n\nReturns a JSON object containing various state info.\n\nArguments:\nNone\n\nResult:\n{\n \"version\": n,          (numeric) The version of the server\n \"protocolversion\": n,  (numeric) The latest supported protocol version\n \"walletversion\": n,    (numeric) The version of the address manager database\n \"balance\": n.nnn,      (numeric) The balance of all accounts calculated with one block confirmation\n \"blocks\": n,           (numeric) The number of blocks processed\n \"timeoffset\": n,       (numeric) The time offset\n \"connections\": n,      (numeric) The number of connected peers\n \"proxy\": \"value\",      (string)  The proxy used by the server\n \"difficulty\": n.nnn,   (numeric) The current target difficulty\n \"testnet\": true|false, (boolean) Whether or not server is using testnet\n \"keypoololdest\": n,    (numeric) Unset\n \"keypoolsize\": n,      (numeric) Unset\n \"unlocked_until\": n,   (numeric) Unset\n \"paytxfee\": n.nnn,     (numeric) The increment used each time more fee is required for an authored transaction\n \"relayfee\": n.nnn,     (numeric) The minimum relay fee for non-free transactions in DUO/KB\n \"errors\": \"value\",     (string)  Any current errors\n}                       \n",
		"getnewaddress":           "getnewaddress (\"account\")\n\nGenerates and returns a new payment address.\n\nArguments:\n1. account (string, optional) DEPRECATED -- Account name the new address will belong to (default=\"default\")\n\nResult:\n\"value\" (string) The payment address\n",
		"getrawchangeaddress":     "getrawchangeaddress (\"account\")\n\nGenerates and returns a new internal payment address for use as a change address in raw transactions.\n\nArguments:\n1. account (string, optional) Account name the new internal address will belong to (default=\"default\")\n\nResult:\n\"value\" (string) The internal payment address\n",
		"getreceivedbyaccount":    "getreceivedbyaccount \"account\" (minconf=1)\n\nDEPRECATED -- Returns the total amount received by addresses of some account, including spent outputs.\n\nArguments:\n1. account (string, required)             Account name to query total received amount for\n2. minconf (numeric, optional, default=1) Minimum number of block confirmations required before an output's value is included in the total\n\nResult:\nn.nnn (numeric) The total received amount valued in bitcoin\n",
		"getreceivedbyaddress":    "getreceivedbyaddress \"address\" (minconf=1)\n\nReturns the total amount received by a single address, including spent outputs.\n\nArguments:\n1. address (string, required)             Payment address which received outputs to include in total\n2. minconf (numeric, optional, default=1) Minimum number of block confirmations required before an output's value is included in the total\n\nResult:\nn.nnn (numeric) The total received amount valued in bitcoin\n",
		"gettransaction":          "gettransaction \"txid\" (includewatchonly=false)\n\nReturns a JSON object with details regarding a transaction relevant to this wallet.\n\nArguments:\n1. txid             (string, required)                 Hash of the transaction to query\n2. includewatchonly (boolean, optional, default=false) Also consider transactions involving watched addresses\n\nResult:\n{\n \"amount\": n.nnn,                  (numeric)         The total amount this transaction credits to the wallet, valued in bitcoin\n \"fee\": n.nnn,                     (numeric)         The total input value minus the total output value, or 0 if 'txid' is not a sent transaction\n \"confirmations\": n,               (numeric)         The number of block confirmations of the transaction\n \"blockhash\": \"value\",             (string)          The hash of the block this transaction is mined in, or the empty string if unmined\n \"blockindex\": n,                  (numeric)         Unset\n \"blocktime\": n,                   (numeric)         The Unix time of the block header this transaction is mined in, or 0 if unmined\n \"txid\": \"value\",                  (string)          The transaction hash\n \"walletconflicts\": [\"value\",...], (array of string) Unset\n \"time\": n,                        (numeric)         The earliest Unix time this transaction was known to exist\n \"timereceived\": n,                (numeric)         The earliest Unix time this transaction was known to exist\n \"details\": [{                     (array of object) Additional details for each recorded wallet credit and debit\n  \"account\": \"value\",              (string)          DEPRECATED -- Unset\n  \"address\": \"value\",              (string)          The address an output was paid to, or the empty string if the output is nonstandard or this detail is regarding a transaction input\n  \"amount\": n.nnn,                 (numeric)         The amount of a received output\n  \"category\": \"value\",             (string)          The kind of detail: \"send\" for sent transactions, \"immature\" for immature coinbase outputs, \"generate\" for mature coinbase outputs, or \"recv\" for all other received outputs\n  \"involveswatchonly\": true|false, (boolean)         Unset\n  \"fee\": n.nnn,                    (numeric)         The included fee for a sent transaction\n  \"vout\": n,                       (numeric)         The transaction output index\n },...],                                             \n \"hex\": \"value\",                   (string)          The transaction encoded as a hexadecimal string\n}                                  \n",
		"help":                    "help (\"command\")\n\nReturns a list of all commands or help for a specified command.\n\nArguments:\n1. command (string, optional) The command to retrieve help for\n\nResult (no command provided):\n\"value\" (string) List of commands\n\nResult (command specified):\n\"value\" (string) Help for specified command\n",
		"importprivkey":           "importprivkey \"privkey\" (\"label\" rescan=true)\n\nImports a WIF-encoded private key to the 'imported' account.\n\nArguments:\n1. privkey (string, required)                The WIF-encoded private key\n2. label   (string, optional)                Unused (must be unset or 'imported')\n3. rescan  (boolean, optional, default=true) Rescan the blockchain (since the genesis block) for outputs controlled by the imported key\n\nResult:\nNothing\n",
		"keypoolrefill":           "keypoolrefill (newsize=100)\n\nDEPRECATED -- This request does nothing since no keypool is maintained.\n\nArguments:\n1. newsize (numeric, optional, default=100) Unused\n\nResult:\nNothing\n",
		"listaccounts":            "listaccounts (minconf=1)\n\nDEPRECATED -- Returns a JSON object of all accounts and their balances.\n\nArguments:\n1. minconf (numeric, optional, default=1) Minimum number of block confirmations required before an unspent output's value is included in the balance\n\nResult:\n{\n \"The account name\": The account balance valued in bitcoin, (object) JSON object with account names as keys and bitcoin amounts as values\n ...\n}\n",
		"listlockunspent":         "listlockunspent\n\nReturns a JSON array of outpoints marked as locked (with lockunspent) for this wallet session.\n\nArguments:\nNone\n\nResult:\n[{\n \"txid\": \"value\", (string)  The transaction hash of the referenced output\n \"vout\": n,       (numeric) The output index of the referenced output\n},...]\n",
		"listreceivedbyaccount":   "listreceivedbyaccount (minconf=1 includeempty=false includewatchonly=false)\n\nDEPRECATED -- Returns a JSON array of objects listing all accounts and the total amount received by each account.\n\nArguments:\n1. minconf          (numeric, optional, default=1)     Minimum number of block confirmations required before a transaction is considered\n2. includeempty     (boolean, optional, default=false) Unused\n3. includewatchonly (boolean, optional, default=false) Unused\n\nResult:\n[{\n \"account\": \"value\", (string)  The name of the account\n \"amount\": n.nnn,    (numeric) Total amount received by payment addresses of the account valued in bitcoin\n \"confirmations\": n, (numeric) Number of block confirmations of the most recent transaction relevant to the account\n},...]\n",
		"listreceivedbyaddress":   "listreceivedbyaddress (minconf=1 includeempty=false includewatchonly=false)\n\nReturns a JSON array of objects listing wallet payment addresses and their total received amounts.\n\nArguments:\n1. minconf          (numeric, optional, default=1)     Minimum number of block confirmations required before a transaction is considered\n2. includeempty     (boolean, optional, default=false) Unused\n3. includewatchonly (boolean, optional, default=false) Unused\n\nResult:\n[{\n \"account\": \"value\",              (string)          DEPRECATED -- Unset\n \"address\": \"value\",              (string)          The payment address\n \"amount\": n.nnn,                 (numeric)         Total amount received by the payment address valued in bitcoin\n \"confirmations\": n,              (numeric)         Number of block confirmations of the most recent transaction relevant to the address\n \"txids\": [\"value\",...],          (array of string) Transaction hashes of all transactions involving this address\n \"involvesWatchonly\": true|false, (boolean)         Unset\n},...]\n",
		"listsinceblock":          "listsinceblock (\"blockhash\" targetconfirmations=1 includewatchonly=false)\n\nReturns a JSON array of objects listing details of all wallet transactions after some block.\n\nArguments:\n1. blockhash           (string, optional)                 Hash of the parent block of the first block to consider transactions from, or unset to list all transactions\n2. targetconfirmations (numeric, optional, default=1)     Minimum number of block confirmations of the last block in the result object.  Must be 1 or greater.  Note: The transactions array in the result object is not affected by this parameter\n3. includewatchonly    (boolean, optional, default=false) Unused\n\nResult:\n{\n \"transactions\": [{                 (array of object) JSON array of objects containing verbose details of the each transaction\n  \"abandoned\": true|false,          (boolean)         Unset\n  \"account\": \"value\",               (string)          DEPRECATED -- Unset\n  \"address\": \"value\",               (string)          Payment address for a transaction output\n  \"amount\": n.nnn,                  (numeric)         The value of the transaction output valued in bitcoin\n  \"bip125-replaceable\": \"value\",    (string)          Unset\n  \"blockhash\": \"value\",             (string)          The hash of the block this transaction is mined in, or the empty string if unmined\n  \"blockindex\": n,                  (numeric)         Unset\n  \"blocktime\": n,                   (numeric)         The Unix time of the block header this transaction is mined in, or 0 if unmined\n  \"category\": \"value\",              (string)          The kind of transaction: \"send\" for sent transactions, \"immature\" for immature coinbase outputs, \"generate\" for mature coinbase outputs, or \"recv\" for all other received outputs.  Note: A single output may be included multiple times under different categories\n  \"confirmations\": n,               (numeric)         The number of block confirmations of the transaction\n  \"fee\": n.nnn,                     (numeric)         The total input value minus the total output value for sent transactions\n  \"generated\": true|false,          (boolean)         Whether the transaction output is a coinbase output\n  \"involveswatchonly\": true|false,  (boolean)         Unset\n  \"time\": n,                        (numeric)         The earliest Unix time this transaction was known to exist\n  \"timereceived\": n,                (numeric)         The earliest Unix time this transaction was known to exist\n  \"trusted\": true|false,            (boolean)         Unset\n  \"txid\": \"value\",                  (string)          The hash of the transaction\n  \"vout\": n,                        (numeric)         The transaction output index\n  \"walletconflicts\": [\"value\",...], (array of string) Unset\n  \"comment\": \"value\",               (string)          Unset\n  \"otheraccount\": \"value\",          (string)          Unset\n },...],                                              \n \"lastblock\": \"value\",              (string)          Hash of the latest-synced block to be used in later calls to listsinceblock\n}                                   \n",
		"listtransactions":        "listtransactions (\"account\" count=10 from=0 includewatchonly=false)\n\nReturns a JSON array of objects containing verbose details for wallet transactions.\n\nArguments:\n1. account          (string, optional)                 DEPRECATED -- Unused (must be unset or \"*\")\n2. count            (numeric, optional, default=10)    Maximum number of transactions to create results from\n3. from             (numeric, optional, default=0)     Number of transactions to skip before results are created\n4. includewatchonly (boolean, optional, default=false) Unused\n\nResult:\n[{\n \"abandoned\": true|false,          (boolean)         Unset\n \"account\": \"value\",               (string)          DEPRECATED -- Unset\n \"address\": \"value\",               (string)          Payment address for a transaction output\n \"amount\": n.nnn,                  (numeric)         The value of the transaction output valued in bitcoin\n \"bip125-replaceable\": \"value\",    (string)          Unset\n \"blockhash\": \"value\",             (string)          The hash of the block this transaction is mined in, or the empty string if unmined\n \"blockindex\": n,                  (numeric)         Unset\n \"blocktime\": n,                   (numeric)         The Unix time of the block header this transaction is mined in, or 0 if unmined\n \"category\": \"value\",              (string)          The kind of transaction: \"send\" for sent transactions, \"immature\" for immature coinbase outputs, \"generate\" for mature coinbase outputs, or \"recv\" for all other received outputs.  Note: A single output may be included multiple times under different categories\n \"confirmations\": n,               (numeric)         The number of block confirmations of the transaction\n \"fee\": n.nnn,                     (numeric)         The total input value minus the total output value for sent transactions\n \"generated\": true|false,          (boolean)         Whether the transaction output is a coinbase output\n \"involveswatchonly\": true|false,  (boolean)         Unset\n \"time\": n,                        (numeric)         The earliest Unix time this transaction was known to exist\n \"timereceived\": n,                (numeric)         The earliest Unix time this transaction was known to exist\n \"trusted\": true|false,            (boolean)         Unset\n \"txid\": \"value\",                  (string)          The hash of the transaction\n \"vout\": n,                        (numeric)         The transaction output index\n \"walletconflicts\": [\"value\",...], (array of string) Unset\n \"comment\": \"value\",               (string)          Unset\n \"otheraccount\": \"value\",          (string)          Unset\n},...]\n",
		"listunspent":             "listunspent (minconf=1 maxconf=9999999 [\"address\",...])\n\nReturns a JSON array of objects representing unlocked unspent outputs controlled by wallet keys.\n\nArguments:\n1. minconf   (numeric, optional, default=1)       Minimum number of block confirmations required before a transaction output is considered\n2. maxconf   (numeric, optional, default=9999999) Maximum number of block confirmations required before a transaction output is excluded\n3. addresses (array of string, optional)          If set, limits the returned details to unspent outputs received by any of these payment addresses\n\nResult:\n{\n \"txid\": \"value\",         (string)  The transaction hash of the referenced output\n \"vout\": n,               (numeric) The output index of the referenced output\n \"address\": \"value\",      (string)  The payment address that received the output\n \"account\": \"value\",      (string)  The account associated with the receiving payment address\n \"scriptPubKey\": \"value\", (string)  The output script encoded as a hexadecimal string\n \"redeemScript\": \"value\", (string)  Unset\n \"amount\": n.nnn,         (numeric) The amount of the output valued in bitcoin\n \"confirmations\": n,      (numeric) The number of block confirmations of the transaction\n \"spendable\": true|false, (boolean) Whether the output is entirely controlled by wallet keys/scripts (false for partially controlled multisig outputs or outputs to watch-only addresses)\n}                         \n",
		"lockunspent":             "lockunspent unlock [{\"txid\":\"value\",\"vout\":n},...]\n\nLocks or unlocks an unspent output.\nLocked outputs are not chosen for transaction inputs of authored transactions and are not included in 'listunspent' results.\nLocked outputs are volatile and are not saved across wallet restarts.\nIf unlock is true and no transaction outputs are specified, all locked outputs are marked unlocked.\n\nArguments:\n1. unlock       (boolean, required)         True to unlock outputs, false to lock\n2. transactions (array of object, required) Transaction outputs to lock or unlock\n[{\n \"txid\": \"value\", (string)  The transaction hash of the referenced output\n \"vout\": n,       (numeric) The output index of the referenced output\n},...]\n\nResult:\ntrue|false (boolean) The boolean 'true'\n",
		"sendfrom":                "sendfrom \"fromaccount\" \"toaddress\" amount (minconf=1 \"comment\" \"commentto\")\n\nDEPRECATED -- Authors, signs, and sends a transaction that outputs some amount to a payment address.\nA change output is automatically included to send extra output value back to the original account.\n\nArguments:\n1. fromaccount (string, required)             Account to pick unspent outputs from\n2. toaddress   (string, required)             Address to pay\n3. amount      (numeric, required)            Amount to send to the payment address valued in bitcoin\n4. minconf     (numeric, optional, default=1) Minimum number of block confirmations required before a transaction output is eligible to be spent\n5. comment     (string, optional)             Unused\n6. commentto   (string, optional)             Unused\n\nResult:\n\"value\" (string) The transaction hash of the sent transaction\n",
		"sendmany":                "sendmany \"fromaccount\" {\"address\":amount,...} (minconf=1 \"comment\")\n\nAuthors, signs, and sends a transaction that outputs to many payment addresses.\nA change output is automatically included to send extra output value back to the original account.\n\nArguments:\n1. fromaccount (string, required) DEPRECATED -- Account to pick unspent outputs from\n2. amounts     (object, required) Pairs of payment addresses and the output amount to pay each\n{\n \"Address to pay\": Amount to send to the payment address valued in bitcoin, (object) JSON object using payment addresses as keys and output amounts valued in bitcoin to send to each address\n ...\n}\n3. minconf (numeric, optional, default=1) Minimum number of block confirmations required before a transaction output is eligible to be spent\n4. comment (string, optional)             Unused\n\nResult:\n\"value\" (string) The transaction hash of the sent transaction\n",
		"sendtoaddress":           "sendtoaddress \"address\" amount (\"comment\" \"commentto\")\n\nAuthors, signs, and sends a transaction that outputs some amount to a payment address.\nUnlike sendfrom, outputs are always chosen from the default account.\nA change output is automatically included to send extra output value back to the original account.\n\nArguments:\n1. address   (string, required)  Address to pay\n2. amount    (numeric, required) Amount to send to the payment address valued in bitcoin\n3. comment   (string, optional)  Unused\n4. commentto (string, optional)  Unused\n\nResult:\n\"value\" (string) The transaction hash of the sent transaction\n",
		"settxfee":                "settxfee amount\n\nModify the increment used each time more fee is required for an authored transaction.\n\nArguments:\n1. amount (numeric, required) The new fee increment valued in bitcoin\n\nResult:\ntrue|false (boolean) The boolean 'true'\n",
		"signmessage":             "signmessage \"address\" \"message\"\n\nSigns a message using the private key of a payment address.\n\nArguments:\n1. address (string, required) Payment address of private key used to sign the message with\n2. message (string, required) Message to sign\n\nResult:\n\"value\" (string) The signed message encoded as a base64 string\n",
		"signrawtransaction":      "signrawtransaction \"rawtx\" ([{\"txid\":\"value\",\"vout\":n,\"scriptpubkey\":\"value\",\"redeemscript\":\"value\"},...] [\"privkey\",...] flags=\"ALL\")\n\nSigns transaction inputs using private keys from this wallet and request.\nThe valid flags options are ALL, NONE, SINGLE, ALL|ANYONECANPAY, NONE|ANYONECANPAY, and SINGLE|ANYONECANPAY.\n\nArguments:\n1. rawtx    (string, required)                Unsigned or partially unsigned transaction to sign encoded as a hexadecimal string\n2. inputs   (array of object, optional)       Additional data regarding inputs that this wallet may not be tracking\n3. privkeys (array of string, optional)       Additional WIF-encoded private keys to use when creating signatures\n4. flags    (string, optional, default=\"ALL\") Sighash flags\n\nResult:\n{\n \"hex\": \"value\",         (string)          The resulting transaction encoded as a hexadecimal string\n \"complete\": true|false, (boolean)         Whether all input signatures have been created\n \"errors\": [{            (array of object) Script verification errors (if exists)\n  \"txid\": \"value\",       (string)          The transaction hash of the referenced previous output\n  \"vout\": n,             (numeric)         The output index of the referenced previous output\n  \"scriptSig\": \"value\",  (string)          The hex-encoded signature script\n  \"sequence\": n,         (numeric)         Script sequence number\n  \"error\": \"value\",      (string)          Verification or signing error related to the input\n },...],                                   \n}                        \n",
		"validateaddress":         "validateaddress \"address\"\n\nVerify that an address is valid.\nExtra details are returned if the address is controlled by this wallet.\nThe following fields are valid only when the address is controlled by this wallet (ismine=true): isscript, pubkey, iscompressed, account, addresses, hex, script, and sigsrequired.\nThe following fields are only valid when address has an associated public key: pubkey, iscompressed.\nThe following fields are only valid when address is a pay-to-script-hash address: addresses, hex, and script.\nIf the address is a multisig address controlled by this wallet, the multisig fields will be left unset if the wallet is locked since the redeem script cannot be decrypted.\n\nArguments:\n1. address (string, required) Address to validate\n\nResult:\n{\n \"isvalid\": true|false,      (boolean)         Whether or not the address is valid\n \"address\": \"value\",         (string)          The payment address (only when isvalid is true)\n \"ismine\": true|false,       (boolean)         Whether this address is controlled by the wallet (only when isvalid is true)\n \"iswatchonly\": true|false,  (boolean)         Unset\n \"isscript\": true|false,     (boolean)         Whether the payment address is a pay-to-script-hash address (only when isvalid is true)\n \"pubkey\": \"value\",          (string)          The associated public key of the payment address, if any (only when isvalid is true)\n \"iscompressed\": true|false, (boolean)         Whether the address was created by hashing a compressed public key, if any (only when isvalid is true)\n \"account\": \"value\",         (string)          The account this payment address belongs to (only when isvalid is true)\n \"addresses\": [\"value\",...], (array of string) All associated payment addresses of the script if address is a multisig address (only when isvalid is true)\n \"hex\": \"value\",             (string)          The redeem script \n \"script\": \"value\",          (string)          The class of redeem script for a multisig address\n \"sigsrequired\": n,          (numeric)         The number of required signatures to redeem outputs to the multisig address\n}                            \n",
		"verifymessage":           "verifymessage \"address\" \"signature\" \"message\"\n\nVerify a message was signed with the associated private key of some address.\n\nArguments:\n1. address   (string, required) Address used to sign message\n2. signature (string, required) The signature to verify\n3. message   (string, required) The message to verify\n\nResult:\ntrue|false (boolean) Whether the message was signed with the private key of 'address'\n",
		"walletlock":              "walletlock\n\nLock the wallet.\n\nArguments:\nNone\n\nResult:\nNothing\n",
		"walletpassphrase":        "walletpassphrase \"passphrase\" timeout\n\nUnlock the wallet.\n\nArguments:\n1. passphrase (string, required)  The wallet passphrase\n2. timeout    (numeric, required) The number of seconds to wait before the wallet automatically locks\n\nResult:\nNothing\n",
		"walletpassphrasechange":  "walletpassphrasechange \"oldpassphrase\" \"newpassphrase\"\n\nChange the wallet passphrase.\n\nArguments:\n1. oldpassphrase (string, required) The old wallet passphrase\n2. newpassphrase (string, required) The new wallet passphrase\n\nResult:\nNothing\n",
		"createnewaccount":        "createnewaccount \"account\"\n\nCreates a new account.\nThe wallet must be unlocked for this request to succeed.\n\nArguments:\n1. account (string, required) Name of the new account\n\nResult:\nNothing\n",
		"exportwatchingwallet":    "exportwatchingwallet (\"account\" download=false)\n\nCreates and returns a duplicate of the wallet database without any private keys to be used as a watching-only wallet.\n\nArguments:\n1. account  (string, optional)                 Unused (must be unset or \"*\")\n2. download (boolean, optional, default=false) Unused\n\nResult:\n\"value\" (string) The watching-only database encoded as a base64 string\n",
		"getbestblock":            "getbestblock\n\nReturns the hash and height of the newest block in the best chain that wallet has finished syncing with.\n\nArguments:\nNone\n\nResult:\n{\n \"hash\": \"value\", (string)  The hash of the block\n \"height\": n,     (numeric) The blockchain height of the block\n}                 \n",
		"getunconfirmedbalance":   "getunconfirmedbalance (\"account\")\n\nCalculates the unspent output value of all unmined transaction outputs for an account.\n\nArguments:\n1. account (string, optional) The account to query the unconfirmed balance for (default=\"default\")\n\nResult:\nn.nnn (numeric) Total amount of all unmined unspent outputs of the account valued in bitcoin.\n",
		"listaddresstransactions": "listaddresstransactions [\"address\",...] (\"account\")\n\nReturns a JSON array of objects containing verbose details for wallet transactions pertaining some addresses.\n\nArguments:\n1. addresses (array of string, required) Addresses to filter transaction results by\n2. account   (string, optional)          Unused (must be unset or \"*\")\n\nResult:\n[{\n \"abandoned\": true|false,          (boolean)         Unset\n \"account\": \"value\",               (string)          DEPRECATED -- Unset\n \"address\": \"value\",               (string)          Payment address for a transaction output\n \"amount\": n.nnn,                  (numeric)         The value of the transaction output valued in bitcoin\n \"bip125-replaceable\": \"value\",    (string)          Unset\n \"blockhash\": \"value\",             (string)          The hash of the block this transaction is mined in, or the empty string if unmined\n \"blockindex\": n,                  (numeric)         Unset\n \"blocktime\": n,                   (numeric)         The Unix time of the block header this transaction is mined in, or 0 if unmined\n \"category\": \"value\",              (string)          The kind of transaction: \"send\" for sent transactions, \"immature\" for immature coinbase outputs, \"generate\" for mature coinbase outputs, or \"recv\" for all other received outputs.  Note: A single output may be included multiple times under different categories\n \"confirmations\": n,               (numeric)         The number of block confirmations of the transaction\n \"fee\": n.nnn,                     (numeric)         The total input value minus the total output value for sent transactions\n \"generated\": true|false,          (boolean)         Whether the transaction output is a coinbase output\n \"involveswatchonly\": true|false,  (boolean)         Unset\n \"time\": n,                        (numeric)         The earliest Unix time this transaction was known to exist\n \"timereceived\": n,                (numeric)         The earliest Unix time this transaction was known to exist\n \"trusted\": true|false,            (boolean)         Unset\n \"txid\": \"value\",                  (string)          The hash of the transaction\n \"vout\": n,                        (numeric)         The transaction output index\n \"walletconflicts\": [\"value\",...], (array of string) Unset\n \"comment\": \"value\",               (string)          Unset\n \"otheraccount\": \"value\",          (string)          Unset\n},...]\n",
		"listalltransactions":     "listalltransactions (\"account\")\n\nReturns a JSON array of objects in the same format as 'listtransactions' without limiting the number of returned objects.\n\nArguments:\n1. account (string, optional) Unused (must be unset or \"*\")\n\nResult:\n[{\n \"abandoned\": true|false,          (boolean)         Unset\n \"account\": \"value\",               (string)          DEPRECATED -- Unset\n \"address\": \"value\",               (string)          Payment address for a transaction output\n \"amount\": n.nnn,                  (numeric)         The value of the transaction output valued in bitcoin\n \"bip125-replaceable\": \"value\",    (string)          Unset\n \"blockhash\": \"value\",             (string)          The hash of the block this transaction is mined in, or the empty string if unmined\n \"blockindex\": n,                  (numeric)         Unset\n \"blocktime\": n,                   (numeric)         The Unix time of the block header this transaction is mined in, or 0 if unmined\n \"category\": \"value\",              (string)          The kind of transaction: \"send\" for sent transactions, \"immature\" for immature coinbase outputs, \"generate\" for mature coinbase outputs, or \"recv\" for all other received outputs.  Note: A single output may be included multiple times under different categories\n \"confirmations\": n,               (numeric)         The number of block confirmations of the transaction\n \"fee\": n.nnn,                     (numeric)         The total input value minus the total output value for sent transactions\n \"generated\": true|false,          (boolean)         Whether the transaction output is a coinbase output\n \"involveswatchonly\": true|false,  (boolean)         Unset\n \"time\": n,                        (numeric)         The earliest Unix time this transaction was known to exist\n \"timereceived\": n,                (numeric)         The earliest Unix time this transaction was known to exist\n \"trusted\": true|false,            (boolean)         Unset\n \"txid\": \"value\",                  (string)          The hash of the transaction\n \"vout\": n,                        (numeric)         The transaction output index\n \"walletconflicts\": [\"value\",...], (array of string) Unset\n \"comment\": \"value\",               (string)          Unset\n \"otheraccount\": \"value\",          (string)          Unset\n},...]\n",
		"renameaccount":           "renameaccount \"oldaccount\" \"newaccount\"\n\nRenames an account.\n\nArguments:\n1. oldaccount (string, required) The old account name to rename\n2. newaccount (string, required) The new name for the account\n\nResult:\nNothing\n",
		"walletislocked":          "walletislocked\n\nReturns whether or not the wallet is locked.\n\nArguments:\nNone\n\nResult:\ntrue|false (boolean) Whether the wallet is locked\n",
	}
}

var LocaleHelpDescs = map[string]func() map[string]string{
	"en_US": HelpDescsEnUS,
}
var RequestUsages = "addmultisigaddress nrequired [\"key\",...] (\"account\")\ncreatemultisig nrequired [\"key\",...]\ndumpprivkey \"address\"\ngetaccount \"address\"\ngetaccountaddress \"account\"\ngetaddressesbyaccount \"account\"\ngetbalance (\"account\" minconf=1)\ngetbestblockhash\ngetblockcount\ngetinfo\ngetnewaddress (\"account\")\ngetrawchangeaddress (\"account\")\ngetreceivedbyaccount \"account\" (minconf=1)\ngetreceivedbyaddress \"address\" (minconf=1)\ngettransaction \"txid\" (includewatchonly=false)\nhelp (\"command\")\nimportprivkey \"privkey\" (\"label\" rescan=true)\nkeypoolrefill (newsize=100)\nlistaccounts (minconf=1)\nlistlockunspent\nlistreceivedbyaccount (minconf=1 includeempty=false includewatchonly=false)\nlistreceivedbyaddress (minconf=1 includeempty=false includewatchonly=false)\nlistsinceblock (\"blockhash\" targetconfirmations=1 includewatchonly=false)\nlisttransactions (\"account\" count=10 from=0 includewatchonly=false)\nlistunspent (minconf=1 maxconf=9999999 [\"address\",...])\nlockunspent unlock [{\"txid\":\"value\",\"vout\":n},...]\nsendfrom \"fromaccount\" \"toaddress\" amount (minconf=1 \"comment\" \"commentto\")\nsendmany \"fromaccount\" {\"address\":amount,...} (minconf=1 \"comment\")\nsendtoaddress \"address\" amount (\"comment\" \"commentto\")\nsettxfee amount\nsignmessage \"address\" \"message\"\nsignrawtransaction \"rawtx\" ([{\"txid\":\"value\",\"vout\":n,\"scriptpubkey\":\"value\",\"redeemscript\":\"value\"},...] [\"privkey\",...] flags=\"ALL\")\nvalidateaddress \"address\"\nverifymessage \"address\" \"signature\" \"message\"\nwalletlock\nwalletpassphrase \"passphrase\" timeout\nwalletpassphrasechange \"oldpassphrase\" \"newpassphrase\"\ncreatenewaccount \"account\"\nexportwatchingwallet (\"account\" download=false)\ngetbestblock\ngetunconfirmedbalance (\"account\")\nlistaddresstransactions [\"address\",...] (\"account\")\nlistalltransactions (\"account\")\nrenameaccount \"oldaccount\" \"newaccount\"\nwalletislocked"
