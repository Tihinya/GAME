package engine

// For using in tests because engine functions will now use socket functions
// and to avoid running into a import cycle it is necessary to use
// these exported functions
//
// !!! All test files must include `package packageName_test` !!!
var (
	ExportTestEntityManager     = entityManager
	ExportTestPositionManager   = positionManager
	ExportTestMotionManager     = motionManager
	ExportTestInputManager      = inputManager
	ExportTestTimerManager      = timerManager
	ExportTestBombManager       = bombManager
	ExportTestExplosionManager  = explosionManager
	ExportTestCollisionManager  = collisionManager
	ExportTestPowerUpManager    = powerUpManager
	ExportTestDamageManager     = damageManager
	ExportTestHealthManager     = healthManager
	ExportTestBoxManager        = boxManager
	ExportTestWallManager       = wallManager
	ExportTestUserEntityManager = userEntityManager
)

var (
	ExportTestMotionSystem    = motionSystem
	ExportTestInputSystem     = inputSystem
	ExportTestPowerUpSystem   = powerUpSystem
	ExportTestHealthSystem    = healthSystem
	ExportTestExplosionSystem = explosionSystem
)
