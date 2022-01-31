#type Lifecycle
type Lifecycle interface{
    Append(Hook)
}

Lifecycle allows constructor to register callbacks that are executed on application start and stop.

#func Options
Options converts a collection of Options into a single Option. This allows packages to bundle sophisticated functionality into easy to use Fx modlule.

#func Invoke
Invoke registers functions that are executed eagerly on application start. Arguments for these invocations are built using the constructors registered by Provide. Passing multiple Invoke options appends the new invocations to the application's existing list.

#fx.New(opts)
New creates and initializes an App, immediately executing any functions registered via invoking options.