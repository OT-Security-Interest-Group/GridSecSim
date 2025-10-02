import random
import time
from opendssdirect import dss

# --- Setup circuit ---
dss.Basic.ClearAll()
dss.Text.Command("Clear")
dss.Text.Command("New Circuit.TransTest basekv=115 pu=1.00 frequency=60")
dss.Text.Command("Edit Vsource.Source bus1=SourceBus.1.2.3 basekv=115 pu=1.0 angle=0 phases=3 frequency=60")

# Transformer
dss.Text.Command("New Transformer.SubXFMR Phases=3 Windings=2 Xhl=7")
dss.Text.Command("~ wdg=1 bus=SourceBus conn=Wye kv=115 kva=25000 %r=0.5")
dss.Text.Command("~ wdg=2 bus=LoadBus conn=Delta kv=13.8 kva=25000 %r=0.5")

# Initial load (will change dynamically)
dss.Text.Command("New Load.DynamicLoad bus1=LoadBus.1.2.3 phases=3 kv=13.8 kw=5000 kvar=1500")

# --- Control loop parameters ---
target_voltage = 13800    # target line-to-line (approx)
step_kw = 500             # step size for load adjustments
iterations = 10

# --- Control loop ---
for i in range(iterations):
    # Random starting or "server-provided" demand
    current_kw = random.randint(3000, 7000)
    dss.Loads.First()  # activate our only load
    dss.Loads.kW(current_kw)

    # Solve
    dss.Text.Command("Solve")

    # Measure load bus voltages
    dss.Circuit.SetActiveBus("LoadBus")
    volts = dss.Bus.VMagAngle()
    avg_voltage = sum(volts[0::2]) / 3  # average of 3 phases

    print(f"\nStep {i+1}: Load set to {current_kw} kW")
    print(f"Average voltage at LoadBus: {avg_voltage:.1f} V")

    # Simple control: step load up or down to move toward target
    if avg_voltage < target_voltage * 0.98:
        new_kw = current_kw - step_kw  # too low voltage, reduce demand
    elif avg_voltage > target_voltage * 1.02:
        new_kw = current_kw + step_kw  # too high voltage, increase demand
    else:
        new_kw = current_kw  # within tolerance

    dss.Loads.kW(new_kw)
    print(f"Adjusted load for next step: {new_kw} kW")

    time.sleep(0.5)  # pacing loop (not needed unless you want "real-time")
