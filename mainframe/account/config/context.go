package config

type ContextKey string

// Main context parameters
const ContextConfig ContextKey = "config"
const ContextAbi ContextKey = "abi"

// Path and query parameters
const ContextAccountId ContextKey = "accountId"

const ContextFrom ContextKey = "from"
const ContextLimit ContextKey = "limit"
