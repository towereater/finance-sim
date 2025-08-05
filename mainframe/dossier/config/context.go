package config

type ContextKey string

// Main context parameters
const ContextConfig ContextKey = "config"
const ContextAbi ContextKey = "abi"
const ContextCab ContextKey = "cab"
const ContextAuth ContextKey = "auth"

// Path and query parameters
const ContextDossier ContextKey = "dossier"

const ContextFrom ContextKey = "from"
const ContextLimit ContextKey = "limit"
