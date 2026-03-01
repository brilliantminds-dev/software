# Stratus

**Stratus** is a minimal, serverless Go web framework for building APIs—designed for production teams leveraging AWS Lambda; audit-ready security logging, and observability integration (traces/metrics/log correlation).

---

## ✨ Key Features

### 🛡️ Middleware-First Pipeline
The request lifecycle is managed through a robust, order-dependent pipeline:
* **Security Headers**: Built-in protection for HSTS, CSP, and XSS.
* **Access Control**: Integrated authentication and rate limiting.
* **Custom Hooks**: Easily inject logic before or after request execution.

**Outcome & Reason**: Explicit logging of why a request was successful or blocked.

### 🔭 First-Class Observability
Designed for distributed systems and modern monitoring stacks:
* **Trace Correlation**: Automatic propagation of Trace IDs across logs and metrics.
* **Vendor Agnostic**: Integrated hooks for **OpenTelemetry**, **New Relic**, and more.
* **Audit Events**: Context-rich logging for deep-dive debugging.
