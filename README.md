# `docker.io/jharper-sec/contrast-security`
The Paketo Buildpack for ContrastSecurity is a Cloud Native Buildpack that contributes the [ContrastSecurity][c] Agent and configures it to connect to the service.

[c]: https://www.contrastsecurity.com

## Behavior
This buildpack will participate if all the following conditions are met

* A binding exists with `type` of `ContrastSecurity`

The buildpack will do the following for Java applications:

* Contributes a Java agent to a layer and configures `$JAVA_TOOL_OPTIONS` to use it
* Contribute external configuration if available
* Transforms the contents of the binding secret to environment variables with the pattern `CONTRAST_<KEY>=<VALUE>`

## Configuration
| Environment Variable          | Description                                     |
| ----------------------------- | ----------------------------------------------- |
| `$CONTRAST_APPLICATION_NAME`  | Configure the ContrastSecurity application name |
| `$CONTRAST_AGENT_SERVER_NAME` | Configure the ContrastSecurity server name      |

## Bindings
The buildpack optionally accepts the following bindings:

### Type: `dependency-mapping`
| Key                   | Value   | Description                                                                                       |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------- |
| `<dependency-digest>` | `<uri>` | If needed, the buildpack will fetch the dependency with digest `<dependency-digest>` from `<uri>` |

## ARM64 Support

This buildpack supports running on ARM64, however, not for all language families. Presently, it supports the Java Agent on ARM64.

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
