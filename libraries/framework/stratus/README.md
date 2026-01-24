# Stratus

**Stratus** is a minimal, secure-by-default Go web framework for building APIs—designed for production teams that care about validation, audit-ready security logging, and first-class observability (traces/metrics/log correlation).

---

## ✨ Key Features

### 🛡️ Middleware-First Pipeline
The request lifecycle is managed through a robust, order-dependent pipeline:
* **Security Headers**: Built-in protection for HSTS, CSP, and XSS.
* **Access Control**: Integrated authentication and rate limiting.
* **Custom Hooks**: Easily inject logic before or after request execution.

### ✅ Request Validation
Achieve consistency across your API with automated validation:
* **Deep Inspection**: Validates request bodies, query parameters, and headers.
* **Standardized Errors**: Ensures every validation failure returns a consistent JSON schema to the client.

### 📝 Security & Audit Logging
Built for compliance and forensics, Stratus captures granular event data:
* **Request ID**: A unique identifier for every transaction.
* **Actor Attribution**: Automatically tracks the user or service account responsible.
* **Outcome & Reason**: Explicit logging of why a request was successful or blocked.

### 🔭 First-Class Observability
Designed for distributed systems and modern monitoring stacks:
* **Trace Correlation**: Automatic propagation of Trace IDs across logs and metrics.
* **Vendor Agnostic**: Integrated hooks for **OpenTelemetry**, **New Relic**, and more.
* **Audit Events**: Context-rich logging for deep-dive debugging.

---

## 📊 Observability Matrix

| Component | Description |
| :--- | :--- |
| **Request ID** | Tracks the lifecycle of a single incoming call. |
| **Actor** | Identifies the identity/service making the request. |
| **Outcome** | Logs the final status (Allow/Deny/Fail). |
| **