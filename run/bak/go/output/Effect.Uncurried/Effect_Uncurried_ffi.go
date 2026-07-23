package Effect_Uncurried

import "gopurs/output/gopurs_runtime"

func MkEffectFn1(f gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(f, a)
		}
	}
}

func MkEffectFn2(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b)
		}
	}
}

func MkEffectFn3(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c)
		}
	}
}

func MkEffectFn4(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d)
		}
	}
}

func MkEffectFn5(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e)
		}
	}
}

func MkEffectFn6(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g)
		}
	}
}

func MkEffectFn7(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g, h gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g), h)
		}
	}
}

func MkEffectFn8(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g, h, i gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g), h), i)
		}
	}
}

func MkEffectFn9(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g, h, i, j gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g), h), i), j)
		}
	}
}

func MkEffectFn10(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g, h, i, j, k gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g), h), i), j), k)
		}
	}
}

func RunEffectFn1(f gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a)
		}
	}
}

func RunEffectFn2(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a, b)
			}
		}
	}
}

func RunEffectFn3(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a, b, c)
				}
			}
		}
	}
}

func RunEffectFn4(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a, b, c, d)
					}
				}
			}
		}
	}
}

func RunEffectFn5(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func() gopurs_runtime.Value {
							return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a, b, c, d, e)
						}
					}
				}
			}
		}
	}
}

func RunEffectFn6(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								var args []gopurs_runtime.Value
								args = append(args, a, b, c, d, e, g)
								return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn7(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func(h gopurs_runtime.Value) func() gopurs_runtime.Value {
								return func() gopurs_runtime.Value {
									var args []gopurs_runtime.Value
									args = append(args, a, b, c, d, e, g, h)
									return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn8(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func(h gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
								return func(i gopurs_runtime.Value) func() gopurs_runtime.Value {
									return func() gopurs_runtime.Value {
										var args []gopurs_runtime.Value
										args = append(args, a, b, c, d, e, g, h, i)
										return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn9(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func(h gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
								return func(i gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
									return func(j gopurs_runtime.Value) func() gopurs_runtime.Value {
										return func() gopurs_runtime.Value {
											var args []gopurs_runtime.Value
											args = append(args, a, b, c, d, e, g, h, i, j)
											return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn10(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func(h gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
								return func(i gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
									return func(j gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
										return func(k gopurs_runtime.Value) func() gopurs_runtime.Value {
											return func() gopurs_runtime.Value {
												var args []gopurs_runtime.Value
												args = append(args, a, b, c, d, e, g, h, i, j, k)
												return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_MkEffectFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn1(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
})
var _Gopurs_MkEffectFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn2(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkEffectFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn3(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkEffectFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn4(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkEffectFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn5(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkEffectFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn6(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkEffectFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn7(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkEffectFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn8(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkEffectFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn9(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkEffectFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn10(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_RunEffectFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn1(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
})
var _Gopurs_RunEffectFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn2(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
})
var _Gopurs_RunEffectFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn3(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
		})
})
var _Gopurs_RunEffectFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn4(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
		})
		})
})
var _Gopurs_RunEffectFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn5(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
		})
		})
		})
})
var _Gopurs_RunEffectFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn6(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
		})
		})
		})
		})
})
var _Gopurs_RunEffectFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn7(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
		})
		})
		})
		})
		})
})
var _Gopurs_RunEffectFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn8(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
		})
		})
		})
		})
		})
		})
})
var _Gopurs_RunEffectFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn9(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
		})
		})
		})
		})
		})
		})
		})
})
var _Gopurs_RunEffectFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn10(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return inner_res
		})
		})
		})
		})
		})
		})
		})
		})
		})
		})
		})
})
