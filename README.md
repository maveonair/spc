# SPC (Solus Package Checker)

This program checks on GitHub if a newer version of a configured program exists (see `config.example.toml`) and exposes the result as a Prometheus Metrics `spc_success_probe`.

## Endpoint

- Prometheus Metrics Endpoint: `:9100/metrics`

## Configuration

### Environment variables

| Variable        | Description          | Default                  |
| --------------- | -------------------- | ------------------------ |
| SPC_CONFIG_FILE | The config file path | `$HOME/.spc/config.toml` |
