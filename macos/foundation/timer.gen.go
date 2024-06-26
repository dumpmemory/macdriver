// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Timer] class.
var TimerClass = _TimerClass{objc.GetClass("NSTimer")}

type _TimerClass struct {
	objc.Class
}

// An interface definition for the [Timer] class.
type ITimer interface {
	objc.IObject
	Invalidate()
	Fire()
	IsValid() bool
	Tolerance() TimeInterval
	SetTolerance(value TimeInterval)
	TimeInterval() TimeInterval
	FireDate() Date
	SetFireDate(value IDate)
	UserInfo() objc.Object
}

// A timer that fires after a certain time interval has elapsed, sending a specified message to a target object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer?language=objc
type Timer struct {
	objc.Object
}

func TimerFrom(ptr unsafe.Pointer) Timer {
	return Timer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ Timer) InitWithFireDateIntervalRepeatsBlock(date IDate, interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.Call[Timer](t_, objc.Sel("initWithFireDate:interval:repeats:block:"), date, interval, repeats, block)
	return rv
}

// Initializes a timer for the specified date and time interval with the specified block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091887-initwithfiredate?language=objc
func NewTimerWithFireDateIntervalRepeatsBlock(date IDate, interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	instance := TimerClass.Alloc().InitWithFireDateIntervalRepeatsBlock(date, interval, repeats, block)
	instance.Autorelease()
	return instance
}

func (t_ Timer) InitWithFireDateIntervalTargetSelectorUserInfoRepeats(date IDate, ti TimeInterval, t objc.IObject, s objc.Selector, ui objc.IObject, rep bool) Timer {
	rv := objc.Call[Timer](t_, objc.Sel("initWithFireDate:interval:target:selector:userInfo:repeats:"), date, ti, t, s, ui, rep)
	return rv
}

// Initializes a timer using the specified object and selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415700-initwithfiredate?language=objc
func NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats(date IDate, ti TimeInterval, t objc.IObject, s objc.Selector, ui objc.IObject, rep bool) Timer {
	instance := TimerClass.Alloc().InitWithFireDateIntervalTargetSelectorUserInfoRepeats(date, ti, t, s, ui, rep)
	instance.Autorelease()
	return instance
}

func (tc _TimerClass) Alloc() Timer {
	rv := objc.Call[Timer](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TimerClass) New() Timer {
	rv := objc.Call[Timer](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTimer() Timer {
	return TimerClass.New()
}

func (t_ Timer) Init() Timer {
	rv := objc.Call[Timer](t_, objc.Sel("init"))
	return rv
}

// Initializes a timer object with the specified time interval and block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091888-timerwithtimeinterval?language=objc
func (tc _TimerClass) TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.Call[Timer](tc, objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}

// Initializes a timer object with the specified time interval and block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091888-timerwithtimeinterval?language=objc
func Timer_TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	return TimerClass.TimerWithTimeIntervalRepeatsBlock(interval, repeats, block)
}

// Initializes a timer object with the specified object and selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1408356-timerwithtimeinterval?language=objc
func (tc _TimerClass) TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	rv := objc.Call[Timer](tc, objc.Sel("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}

// Initializes a timer object with the specified object and selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1408356-timerwithtimeinterval?language=objc
func Timer_TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	return TimerClass.TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti, aTarget, aSelector, userInfo, yesOrNo)
}

// Initializes a timer object with the specified invocation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1407170-timerwithtimeinterval?language=objc
func (tc _TimerClass) TimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation IInvocation, yesOrNo bool) Timer {
	rv := objc.Call[Timer](tc, objc.Sel("timerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}

// Initializes a timer object with the specified invocation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1407170-timerwithtimeinterval?language=objc
func Timer_TimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation IInvocation, yesOrNo bool) Timer {
	return TimerClass.TimerWithTimeIntervalInvocationRepeats(ti, invocation, yesOrNo)
}

// Creates a new timer and schedules it on the current run loop in the default mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415941-scheduledtimerwithtimeinterval?language=objc
func (tc _TimerClass) ScheduledTimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation IInvocation, yesOrNo bool) Timer {
	rv := objc.Call[Timer](tc, objc.Sel("scheduledTimerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return rv
}

// Creates a new timer and schedules it on the current run loop in the default mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415941-scheduledtimerwithtimeinterval?language=objc
func Timer_ScheduledTimerWithTimeIntervalInvocationRepeats(ti TimeInterval, invocation IInvocation, yesOrNo bool) Timer {
	return TimerClass.ScheduledTimerWithTimeIntervalInvocationRepeats(ti, invocation, yesOrNo)
}

// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1412416-scheduledtimerwithtimeinterval?language=objc
func (tc _TimerClass) ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	rv := objc.Call[Timer](tc, objc.Sel("scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}

// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1412416-scheduledtimerwithtimeinterval?language=objc
func Timer_ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	return TimerClass.ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti, aTarget, aSelector, userInfo, yesOrNo)
}

// Stops the timer from ever firing again and requests its removal from its run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415405-invalidate?language=objc
func (t_ Timer) Invalidate() {
	objc.Call[objc.Void](t_, objc.Sel("invalidate"))
}

// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091889-scheduledtimerwithtimeinterval?language=objc
func (tc _TimerClass) ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.Call[Timer](tc, objc.Sel("scheduledTimerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}

// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091889-scheduledtimerwithtimeinterval?language=objc
func Timer_ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	return TimerClass.ScheduledTimerWithTimeIntervalRepeatsBlock(interval, repeats, block)
}

// Causes the timer's message to be sent to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1414035-fire?language=objc
func (t_ Timer) Fire() {
	objc.Call[objc.Void](t_, objc.Sel("fire"))
}

// A Boolean value that indicates whether the timer is currently valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1408249-valid?language=objc
func (t_ Timer) IsValid() bool {
	rv := objc.Call[bool](t_, objc.Sel("isValid"))
	return rv
}

// The amount of time after the scheduled fire date that the timer may fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415085-tolerance?language=objc
func (t_ Timer) Tolerance() TimeInterval {
	rv := objc.Call[TimeInterval](t_, objc.Sel("tolerance"))
	return rv
}

// The amount of time after the scheduled fire date that the timer may fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415085-tolerance?language=objc
func (t_ Timer) SetTolerance(value TimeInterval) {
	objc.Call[objc.Void](t_, objc.Sel("setTolerance:"), value)
}

// The timer’s time interval, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1409024-timeinterval?language=objc
func (t_ Timer) TimeInterval() TimeInterval {
	rv := objc.Call[TimeInterval](t_, objc.Sel("timeInterval"))
	return rv
}

// The date at which the timer will fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1407353-firedate?language=objc
func (t_ Timer) FireDate() Date {
	rv := objc.Call[Date](t_, objc.Sel("fireDate"))
	return rv
}

// The date at which the timer will fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1407353-firedate?language=objc
func (t_ Timer) SetFireDate(value IDate) {
	objc.Call[objc.Void](t_, objc.Sel("setFireDate:"), value)
}

// The receiver's userInfo object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1408911-userinfo?language=objc
func (t_ Timer) UserInfo() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("userInfo"))
	return rv
}
