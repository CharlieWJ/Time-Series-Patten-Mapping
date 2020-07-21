features = {
    "one" : {
        "neutral_f" : "1.0",
        "min_f" : "1.0",
        "max_f" : "1.0",
        "phi_f" : "math.Max", #Was originally "max"
        "delta_i_f" : "0.0"
    },
    "width" : {
        "neutral_f" : "0.0",
        "min_f" : "0.0",
        "max_f" : "n",
        "phi_f" : "+",
        "delta_i_f" : "1.0"
    },
    "surface" : {
        "neutral_f" : "0.0",
        "min_f" : "math.Inf(-1)",#Was originally "-inf"
        "max_f" : "math.Inf(1)",#Was originally "+inf"
        "phi_f" : "+",
        "delta_i_f" : "xi"
    },
    "max" : {
        "neutral_f" : "math.Inf(-1)",#Was originally "-inf"
        "min_f" : "math.Inf(-1)",#Was originally "-inf"
        "max_f" : "math.Inf(1)",#Was originally "+inf"
        "phi_f" : "math.Max",#Was originally "max" 
        "delta_i_f" : "xi"
    },
    "min" : {
        "neutral_f" : "math.Inf(1)",#Was originally "+inf"
        "min_f" : "math.Inf(-1)",#Was originally "-inf"
        "max_f" : "math.Inf(1)",#Was originally "+inf"
        "phi_f" : "math.Min",#Was originally "min"
        "delta_i_f" : "xi"
    },
    "range" : {
        "neutral_f" : "0.0 //neutral_f",#This only alters D's values.
        "min_f" : "0.0 //min_f",#This alters C's R's & H's values.
        "max_f" : "math.Inf(1) //max_f",#Was originally "+inf"
        "phi_f" : "diff",
        "delta_i_f" : "xi"
    }
 #,
 #   "range" : {
 #       "neutral_f" : "0.0",
 #       "min_f" : "0.0",
 #       "max_f" : "math.Inf(1)",#Was originally "+inf"
 #       "phi_f" : "",
 #       "delta_i_f" : "xi"
 #   }
}