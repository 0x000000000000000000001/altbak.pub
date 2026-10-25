"""Command-local Go compilation options shared by the benchmark runners."""


def go_build_flags(dev):
    # Match generated packages, including their entry point, while leaving the
    # standard library, external modules and the gopurs compiler optimized.
    return ['-gcflags=gopurs/output/...=-N -l'] if dev else []
